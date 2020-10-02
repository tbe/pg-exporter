package generator

import (
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
	"k8s.io/klog"
)

// extractBoolTagOrDie gets the comment-tags for the key and asserts that, if
// it exists, the value is boolean.  If the tag did not exist, it returns
// false.
func extractTag(key string, lines []string) (bool, bool) {
	tag, exists := types.ExtractCommentTags("+", lines)["metric"]
	if !exists {
		return false, false
	}
	if tag[0] == "slice" {
		return true, true
	}
	return true, false
}

func Packages(_ *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	return generator.Packages{&generator.DefaultPackage{
		PackageName: "models",
		PackagePath: arguments.OutputPackagePath,
		HeaderText:  []byte("/*generated by binding gen*/\n"),
		PackageDocumentation: []byte(
			`// Package mapper has auto-generated orm metric mappers.
		`),
		// GeneratorFunc returns a list of generators. Each generator makes a
		// single file.
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			generators = []generator.Generator{
			}
			// Since we want a file per type that we generate a set for, we
			// have to provide a function for this.
			for _, t := range c.Order {
				klog.Infof("considering %v", t)
				generators = append(generators, &mappingGen{
					DefaultGen: generator.DefaultGen{
						// Use the privatized version of the
						// type name as the file name.
						//
						// TODO: make a namer that converts
						// camelCase to '-' separation for file
						// names?
						OptionalName: c.Namers["private"].Name(t) + "_gen",
					},
					outputPackage: arguments.OutputPackagePath,
					typeToMatch:   t,
					imports:       generator.NewImportTracker(),
				})
			}
			return generators
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			// It would be reasonable to filter by the type's package here.
			// It might be necessary if your input directory has a big
			// import graph.
			switch t.Kind {
			case types.Struct:
				enabled, _ := extractTag("metric", t.CommentLines)
				return enabled
			}
			return false
		},
	}}
}