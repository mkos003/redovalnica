// Uporaba ukazne vrstice
// go run main.go --stOcen 3 --minOcena 1 --maxOcena 10
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mkos003/redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {

	// Uporaba .\main.exe --help
	cmd := &cli.Command{
		Name:  "Redovalnica CLI",
		Usage: "Upravljanje študentskih ocen preko ukazne vrstice",

		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "stOcen",
				Usage:   "Minimalno število ocen za pozitivno oceno",
				Value:   3, // default
				Aliases: []string{"s"},
			},
			&cli.IntFlag{
				Name:    "minOcena",
				Usage:   "Najmanjša dovoljena ocena",
				Value:   1,
				Aliases: []string{"min"},
			},
			&cli.IntFlag{
				Name:    "maxOcena",
				Usage:   "Največja dovoljena ocena",
				Value:   10,
				Aliases: []string{"max"},
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {

			// Preberemo vrednosti iz CLI
			// Te vrednosti bi uporabili, če bi studente brali iz uporabniskega vnosa (za postavitev omejitev pri branju)
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			fmt.Printf("Nastavitve: stOcen=%d, minOcena=%d, maxOcena=%d\n\n", stOcen, minOcena, maxOcena)

			// Pripravimo studente
			studenti := make(map[string]redovalnica.Student)
			studenti["1234"] = redovalnica.NewStudent("Janez", "Novak", []int{10, 6, 5})
			studenti["4321"] = redovalnica.NewStudent("Ivan", "Novak", []int{8, 7, 6})
			studenti["5678"] = redovalnica.NewStudent("Petra", "Novak", []int{5, 5, 6})
			studenti["8765"] = redovalnica.NewStudent("Marko", "Novak", []int{9, 10, 9})

			// Ustvarimo redovalnico
			r := redovalnica.NewRedovalnica(studenti)

			r.IzpisVsehOcen()
			fmt.Println()
			r.DodajOceno("1234", 10)
			r.DodajOceno("1234", 9)
			r.IzpisVsehOcen()
			fmt.Println()
			r.IzpisiKoncniUspeh()

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
