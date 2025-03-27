package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	"github.com/kong/kubernetes-ingress-controller/v3/internal/cmd/rootcmd/config"
)

// This program generates markdown documentation for the KIC executable's flags.
// It skips deprecated and hidden flags.
func main() {
	cfg := config.NewCLIConfig()

	markdown := strings.Builder{}
	markdown.WriteString("<!-- This document is generated by KIC's 'generate.docs' make target, DO NOT EDIT -->\n\n")
	markdown.WriteString("## Flags\n\n")
	markdown.WriteString("| Flag | Type | Description | Default |\n")
	markdown.WriteString("| ---- | ---- | ----------- | ------- |\n")

	cfg.FlagSet().VisitAll(func(flag *pflag.Flag) {
		// Skip deprecated and hidden flags.
		if flag.Deprecated != "" || flag.Hidden {
			return
		}

		name := fmt.Sprintf("`--%s`", flag.Name)
		typ := fmt.Sprintf("`%s`", getTypeForHuman(flag))

		description := flag.Usage

		// In case of panic adjust in KIC code the description of the flag.
		// Make sure the first letter is capitalized.
		if first := description[0]; first >= 'a' && first <= 'z' {
			panic(fmt.Sprintf("flag's: %s description must start with a capital letter: %s", flag.Name, description))
		}
		// Make sure the last character is a period.
		if !strings.HasSuffix(description, ".") {
			panic(fmt.Sprintf("flag's: %s description must end with a period: %s", flag.Name, description))
		}

		defaultValue := ""
		if flag.DefValue != "" {
			defaultValue = fmt.Sprintf("`%s`", flag.DefValue)
		}

		markdown.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", name, typ, description, defaultValue))
	})

	fmt.Print(markdown.String())
}

// getTypeForHuman returns a human-friendly type name for the given flag that
// is compatible with native way of printing types in --help output.
// See https://github.com/spf13/pflag/blob/d5e0c0615acee7028e1e2740a11102313be88de1/flag.go#L568-L610.
// Panics if the type is unknown - this way it validates whether type is human readable.
func getTypeForHuman(flag *pflag.Flag) string {
	switch typ := flag.Value.Type(); typ {
	case "float64", "float32":
		return "float"
	case "int64":
		return "int"
	case "uint64":
		return "uint"
	case "stringSlice":
		return "strings"
	case "intSlice":
		return "ints"
	case "uintSlice":
		return "uints"
	case "boolSlice":
		return "bools"
	case "mapStringBool":
		return "list of string=bool"
	case "config.MetricsAccessFilter":
		return "string"
	// The below are types that are human readable out-of-the-box, in case of missing one extend the list.
	case "bool", "string", "int", "uint", "duration", "namespaced-name":
		return typ
	default:
		panic(fmt.Sprintf("unknown type %q", typ))
	}
}
