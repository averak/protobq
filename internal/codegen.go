package internal

import (
	"fmt"
	"runtime/debug"
	"slices"
	"strings"

	"github.com/averak/protobq/internal/protobuf/protobq"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

var (
	timeIdents = struct {
		Minute protogen.GoIdent
	}{
		Minute: protogen.GoImportPath("time").Ident("Minute"),
	}
	protobqIdents = struct {
		MaterializedView        protogen.GoIdent
		MaterializedViewOptions protogen.GoIdent
		InsertDTO               protogen.GoIdent
		internal                struct {
			NewInsertDTOImpl protogen.GoIdent
			BQField          protogen.GoIdent
		}
	}{
		MaterializedView:        protogen.GoImportPath("github.com/averak/protobq").Ident("MaterializedView"),
		MaterializedViewOptions: protogen.GoImportPath("github.com/averak/protobq").Ident("MaterializedViewOptions"),
		InsertDTO:               protogen.GoImportPath("github.com/averak/protobq").Ident("InsertDTO"),
		internal: struct {
			NewInsertDTOImpl protogen.GoIdent
			BQField          protogen.GoIdent
		}{
			NewInsertDTOImpl: protogen.GoImportPath("github.com/averak/protobq/internal").Ident("NewInsertDTOImpl"),
			BQField:          protogen.GoImportPath("github.com/averak/protobq/internal").Ident("BQField"),
		},
	}
)

type CodeGenerator struct {
	plugin *protogen.Plugin
	file   *protogen.File
}

func NewCodeGenerator(plugin *protogen.Plugin, file *protogen.File) *CodeGenerator {
	return &CodeGenerator{
		plugin: plugin,
		file:   file,
	}
}

func (g CodeGenerator) Gen() error {
	if !g.shouldGenerate() {
		return nil
	}

	filename := g.file.GeneratedFilenamePrefix + ".protobq.go"
	gf := g.plugin.NewGeneratedFile(filename, g.file.GoImportPath)

	{ // generate file header
		info, _ := debug.ReadBuildInfo()
		gf.P("// Code generated by ", info.Path, ". DO NOT EDIT.")
		gf.P("// source: ", g.file.Desc.Path())
		gf.P()
		gf.P("package ", g.file.GoPackageName)
		gf.P()
	}
	{ // generate materialized view schema
		for _, msg := range g.file.Messages {
			if !g.isMaterializedViewSchema(msg) {
				continue
			}
			ext, _ := proto.GetExtension(msg.Desc.Options(), protobq.E_MaterializedView).(*protobq.MaterializedView)

			gf.P("var _ ", protobqIdents.MaterializedView, " = (*", msg.GoIdent.GoName, ")(nil)")
			gf.P()

			gf.P("func (mv *", msg.GoIdent.GoName, ") Name() string {")
			gf.P("    return \"", msg.Desc.Name(), "\"")
			gf.P("}")
			gf.P()

			gf.P("func (mv *", msg.GoIdent.GoName, ") Options() ", protobqIdents.MaterializedViewOptions, " {")
			gf.P("    return ", protobqIdents.MaterializedViewOptions, "{")
			gf.P("        EnableRefresh: ", ext.GetEnableRefresh(), ",")
			gf.P("        RefreshInterval: ", ext.GetRefreshIntervalMinutes(), " * ", timeIdents.Minute, ",")
			gf.P("    }")
			gf.P("}")
			gf.P()

			gf.P("func (mv *", msg.GoIdent.GoName, ") InsertDTO() ", protobqIdents.InsertDTO, " {")
			gf.P("    res := ", protobqIdents.internal.NewInsertDTOImpl, "(\"", ext.GetBaseTable(), "\", nil)")
			for _, field := range msg.Fields {
				g.generateAddField(gf, field, nil, "res", "mv")
			}
			gf.P("    return res")
			gf.P("}")
			gf.P()

			//for _, field := range msg.Fields {
			//fieldExt, _ := proto.GetExtension(field.Desc.Options(), protobq.E_MaterializedViewField).(*protobq.MaterializedViewField)
			//if fieldExt != nil {
			//	gf.P("    res[\"", field.Desc.Name(), "\"] = mv.", field.GoName)
			//}
		}
	}
	return nil
}

func (g CodeGenerator) shouldGenerate() bool {
	for _, msg := range g.file.Messages {
		if g.isMaterializedViewSchema(msg) {
			return true
		}
	}
	return false
}

func (g CodeGenerator) isMaterializedViewSchema(msg *protogen.Message) bool {
	opts := msg.Desc.Options()
	if opts == nil {
		return false
	}

	ext, _ := proto.GetExtension(opts, protobq.E_MaterializedView).(*protobq.MaterializedView)
	return ext != nil
}

func (g CodeGenerator) generateAddField(gf *protogen.GeneratedFile, field *protogen.Field, parentPaths []string, result string, receiver string) {
	ext, _ := proto.GetExtension(field.Desc.Options(), protobq.E_MaterializedViewField).(*protobq.MaterializedViewField)
	if ext == nil {
		ext = &protobq.MaterializedViewField{}
	}

	paths := parentPaths
	if len(ext.GetOriginPath()) > 0 {
		paths = append(paths, ext.GetOriginPath()...)
	} else {
		paths = append(paths, string(field.Desc.Name()))
	}

	blacklist := []string{
		"google.protobuf.Timestamp",
	}
	if field.Message != nil && !slices.Contains(blacklist, string(field.Message.Desc.FullName())) {
		for _, nestedField := range field.Message.Fields {
			g.generateAddField(gf, nestedField, paths, result, receiver+".Get"+field.GoName+"()")
		}
	} else {
		gf.P(result, ".AddField(", protobqIdents.internal.BQField, "{[]string{", fmt.Sprintf(`"%s"`, strings.Join(paths, `", "`)), "}, ", receiver, ".Get", field.GoName, "()})")
	}
}
