package cmd

import (
	"felloe/helpers"
	"felloe/js"
	"felloe/js/compiler"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use: "template [script]",
	Short: "Run a script and display output without deploying",
	Args: cobra.MinimumNArgs(1),
	RunE: template,

}

func init() {
	rootCmd.AddCommand(templateCmd)
}

var defaultFn func() []helpers.GenericManifest

func template(cmd *cobra.Command, args []string) error {

	compiled, err := compiler.CompileScript(args[0])
	if err != nil {
		return err
	}

	vm := js.GetRuntime()

	_, err = vm.RunScript(args[0], compiled)
	if err != nil {
		return err
	}

	export := vm.Get("exports").ToObject(vm).Get("default")
	if err = vm.ExportTo(export, &defaultFn); err != nil {
		return err
	}

	yamlOutput := ""
	res := defaultFn()

	sorted := helpers.SortManifests(res)

	for _, k8sObject := range sorted {
		y, err := yaml.Marshal(&k8sObject)
		if err != nil {
			return err
		}

		yamlOutput = yamlOutput + "---\n" + string(y)
	}

	fmt.Println(yamlOutput)

	return nil
}