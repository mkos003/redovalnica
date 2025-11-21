## Redovalnica

`redovalnica` je Go paket za upravljanje ocen dijakov.

### Funkcije
- dodajanje ocen,
- izpis vseh ocen,
- izračun končnega uspeha.

### Primer
```go
studenti := make(map[string]redovalnica.Student)
studenti["1234"] = redovalnica.NewStudent("Janez", "Novak", []int{10, 6, 5})

r := redovalnica.NewRedovalnica(studenti)
r.IzpisVsehOcen()
r.IzpisiKoncniUspeh()