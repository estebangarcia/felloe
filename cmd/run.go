package cmd

import (
	"felloe/compiler"
	"felloe/js"
	"felloe/js/loader"
	"felloe/logger"
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var runCmd = &cobra.Command{
	Use: "run [script]",
	Short: "Run a script",
	Args: cobra.MinimumNArgs(1),
	RunE: run,

}

func init() {
	rootCmd.AddCommand(runCmd)
}

var defaultFn func() []interface{}

func run(cmd *cobra.Command, args []string) error {
	fileContent, err := ioutil.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("couldn't open %v", args[0])
	}

	c, err := compiler.New()
	if err != nil {
		return err
	}

	compiled, err := c.Transform(string(fileContent))
	if err != nil {
		return err
	}

	registry := require.NewRegistryWithLoader(loader.Load)
	loader.InitNativeModules(registry)

	vm := goja.New()
	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
	vm.Set("console", js.NewJSConsole(logger.GetLogger()))
	registry.Enable(vm)

	exports := vm.NewObject()
	_ = vm.Set("exports", exports)

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

	for _, k8sObject := range res {
		y, err := yaml.Marshal(&k8sObject)
		if err != nil {
			return err
		}

		yamlOutput = yamlOutput + "---\n" + string(y)
	}

	fmt.Println(yamlOutput)





	return nil
}