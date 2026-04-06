// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"sort"
	"testing"

	"github.com/google/subcommands"
)

func TestRegisterCommands(t *testing.T) {
	cdr := subcommands.NewCommander(flag.NewFlagSet("test", flag.ContinueOnError), "test")
	registerCommands(cdr)

	cmds := registeredCommands(cdr)

	expected := []string{"commands", "flags", "help", "check", "diff", "gen", "show"}
	for _, name := range expected {
		if !cmds[name] {
			t.Errorf("expected command %q to be registered, but it was not", name)
		}
	}

	if len(cmds) != len(expected) {
		got := make([]string, 0, len(cmds))
		for name := range cmds {
			got = append(got, name)
		}
		sort.Strings(got)
		sort.Strings(expected)
		t.Errorf("expected commands %v, got %v", expected, got)
	}
}

func TestRegisteredCommandsMatchesVisitCommands(t *testing.T) {
	// Verify that registeredCommands returns the same set that VisitCommands
	// reports, ensuring no command is missed or duplicated.
	cdr := subcommands.NewCommander(flag.NewFlagSet("test", flag.ContinueOnError), "test")
	registerCommands(cdr)

	fromVisit := map[string]bool{}
	cdr.VisitCommands(func(_ *subcommands.CommandGroup, cmd subcommands.Command) {
		fromVisit[cmd.Name()] = true
	})

	fromFunc := registeredCommands(cdr)

	for name := range fromVisit {
		if !fromFunc[name] {
			t.Errorf("VisitCommands reported %q but registeredCommands did not", name)
		}
	}
	for name := range fromFunc {
		if !fromVisit[name] {
			t.Errorf("registeredCommands reported %q but VisitCommands did not", name)
		}
	}
}
