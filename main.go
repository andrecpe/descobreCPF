package main

import (
	"flag"
	"fmt"
	"github.com/olekukonko/ts"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	cpfInt []int
	cpfX   []int
	tab    int
	termo  int // 0 -> Testa para CPF e NIS - 1 -> Só para CPF - 2 -> Só para NIS
)

type retorno struct {
	idx int
	val int
}

func init() {
	var cpf string
	flag.StringVar(&cpf, "cpf", "", "`Número do CPF` suposto. Mínimo 8 e máximo 11 dígitos.\nExs.: 12345678901 (Com 11 dígitos só testa a validade), 1234567890 (10 dig), 123456789 (9 dig) ou 12345678 (8 dig).")
	flag.IntVar(&termo, "T", 0, "Teste apenas CPF (1), apenas PIS (2), ou ambos (0 - padrão)")
	flag.Parse()
	if cpf == "" {
		log.SetFlags(0)
		log.SetPrefix("  ")
		log.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
		log.Println("┃                                                                 ┃")
		log.Println("┃  Programa valida ou descobre CPF e NIS com dígitos incompletos  ┃")
		log.Println("┃                                                                 ┃")
		log.Println("┃                                     Desenvolvedor: André Costa  ┃")
		log.Println("┃                                             andrecpe@gmail.com  ┃")
		log.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
		nome0 := strings.Split(os.Args[0], `\`)
		nome := strings.Split(nome0[len(nome0)-1], ".")
		log.Println()
		log.Println("Aceita números incompletos, ou X no lugar do dígito que falta.")
		log.Println("Exemplo de uso:", nome[0], "-cpf 123456789")
		log.Println("Exemplo de uso:", nome[0], "-cpf 12345X67X89")
		log.Println()
		flag.PrintDefaults()
		log.Println()
		os.Exit(1)
	} else {
		key := regexp.MustCompile(`\d|x|X`)
		for i, s := range key.FindAllString(cpf, -1) {
			if s == "X" || s == "x" {
				cpfX = append(cpfX, i)
				cpfInt = append(cpfInt, 0)
			} else {
				tmp, _ := strconv.Atoi(s)
				cpfInt = append(cpfInt, tmp)
			}
			if i > 9 {
				break
			}
		}
	}
}

func main() {
	cpf := cpfInt
	cpfx := cpfX
	var num [11]int
	c := len(cpf)
	x := len(cpfx)
	if c == 11 {
		for i := 0; i < 11; i++ {
			num[i] = cpf[i]
		}
	}
	if x > 0 && c != 11 {
		log.Fatalln("Digite pelo menos 11 caracteres.")
	}
	if x == 0 && c < 8 {
		log.Fatalln("Digite pelo menos 8 caracteres.")
	}
	if x > 5 {
		log.Fatalln("Digite pelo menos 6 caracteres fixos.")
	}

	if x == 0 {
		if c == 11 {
			testesimples(num)
		} else {
			variavel(cpf)
		}
	} else {
		fixo(num, cpfX)
	}
}

func testesimples(num [11]int) {
	var cpf, nis string

	if !isCPF(num) {
		cpf = "IN"
	}
	if !isNIS(num) {
		nis = "IN"
	}
	fmt.Printf("\n   * VALIDAÇÃO DE NÚMERO PIS E CPF *\n\n")
	fmt.Printf("  Para número CPF:\n")
	fmt.Printf("   - %s: %sVÁLIDO\n\n", cpfToFormated(cpfInt), cpf)
	fmt.Printf("  Para número NIS:\n")
	fmt.Printf("   - %s: %sVÁLIDO\n\n", nisToFormated(cpfInt), nis)
}
func variavel(cpfInt []int) {
	var fim [][11]int
	tam := 11
	for a := 0; a < tam; a++ {
		for b := a + 1; b < tam; b++ {
			for c := b + 1; c < tam; c++ {
				for d := c + 1; d < tam; d++ {
					for e := d + 1; e < tam; e++ {
						for f := e + 1; f < tam; f++ {
							for g := f + 1; g < tam; g++ {
								for h := g + 1; h < tam; h++ {
									if len(cpfInt) == 8 {
										add(&fim, retorno{idx: a, val: cpfInt[0]}, retorno{idx: b, val: cpfInt[1]}, retorno{idx: c, val: cpfInt[2]}, retorno{idx: d, val: cpfInt[3]}, retorno{idx: e, val: cpfInt[4]}, retorno{idx: f, val: cpfInt[5]}, retorno{idx: g, val: cpfInt[6]}, retorno{idx: h, val: cpfInt[7]})
									} else {
										for i := h + 1; i < tam; i++ {
											if len(cpfInt) == 9 {
												add(&fim, retorno{idx: a, val: cpfInt[0]}, retorno{idx: b, val: cpfInt[1]}, retorno{idx: c, val: cpfInt[2]}, retorno{idx: d, val: cpfInt[3]}, retorno{idx: e, val: cpfInt[4]}, retorno{idx: f, val: cpfInt[5]}, retorno{idx: g, val: cpfInt[6]}, retorno{idx: h, val: cpfInt[7]}, retorno{idx: i, val: cpfInt[8]})
											} else {
												for j := i + 1; j < tam; j++ {
													add(&fim, retorno{idx: a, val: cpfInt[0]}, retorno{idx: b, val: cpfInt[1]}, retorno{idx: c, val: cpfInt[2]}, retorno{idx: d, val: cpfInt[3]}, retorno{idx: e, val: cpfInt[4]}, retorno{idx: f, val: cpfInt[5]}, retorno{idx: g, val: cpfInt[6]}, retorno{idx: h, val: cpfInt[7]}, retorno{idx: i, val: cpfInt[8]}, retorno{idx: j, val: cpfInt[9]})
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	resultado(fim)
}
func fixo(cpf [11]int, cpfx []int) {
	qt := len(cpfx)
	var res [][11]int
	for i := 0; i < 10; i++ {
		cpf[cpfx[0]] = i
		if qt == 1 {
			res = append(res, cpf)
		} else {
			for j := 0; j < 10; j++ {
				cpf[cpfx[1]] = j
				if qt == 2 {
					res = append(res, cpf)
				} else {
					for k := 0; k < 10; k++ {
						cpf[cpfx[2]] = k
						if qt == 3 {
							res = append(res, cpf)
						} else {
							for l := 0; l < 10; l++ {
								cpf[cpfx[3]] = l
								if qt == 4 {
									res = append(res, cpf)
								} else {
									for m := 0; m < 10; m++ {
										cpf[cpfx[4]] = m
										if qt == 5 {
											res = append(res, cpf)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	resultado(res)
}

func isCPF(cpf [11]int) bool {
	var sum int
	for idx, factor := 0, 10; idx < 9; idx, factor = idx+1, factor-1 {
		sum += cpf[idx] * factor
	}

	sum = (sum * 10) % 11
	if sum == 10 {
		sum = 0
	}
	if sum != cpf[9] {
		return false
	}

	sum = 0
	for idx, factor := 0, 11; idx < 10; idx, factor = idx+1, factor-1 {
		sum += cpf[idx] * factor
	}

	sum = (sum * 10) % 11
	if sum == 10 {
		sum = 0
	}
	if sum != cpf[10] {
		return false
	}
	return true
}

func isNIS(nis [11]int) bool {
	factor := [10]int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var sum int
	for i := 0; i < 10; i++ {
		sum += nis[i] * factor[i]
	}
	sum %= 11
	if sum > 1 {
		sum = 11 - sum
	} else {
		sum = 0
	}
	if sum == nis[10] {
		return true
	} else {
		return false
	}
}

func cpfToFormated(cpf interface{}) (str string) {
	switch cpf.(type) {
	case string:
		str = cpf.(string)
		str = str[0:3] + "." + str[3:6] + "." + str[6:9] + "-" + str[9:11]
	case []int:
		for n, v := range cpf.([]int) {
			t := strconv.Itoa(v)
			str += t
			if n == 2 || n == 5 {
				str += "."
			}
			if n == 8 {
				str += "-"
			}
		}
	case [11]int:
		for n, v := range cpf.([11]int) {
			t := strconv.Itoa(v)
			str += t
			if n == 2 || n == 5 {
				str += "."
			}
			if n == 8 {
				str += "-"
			}
		}
	default:
		log.Fatalln("Tipo não definido.")
	}
	return
}
func nisToFormated(nis interface{}) (str string) {
	switch nis.(type) {
	case string:
		str = nis.(string)
		str = str[0:3] + "." + str[3:8] + "." + str[8:10] + "-" + str[10:11]
	case []int:
		for n, v := range nis.([]int) {
			t := strconv.Itoa(v)
			str += t
			if n == 2 || n == 7 {
				str += "."
			}
			if n == 9 {
				str += "-"
			}
		}
	case [11]int:
		for n, v := range nis.([11]int) {
			t := strconv.Itoa(v)
			str += t
			if n == 2 || n == 7 {
				str += "."
			}
			if n == 9 {
				str += "-"
			}
		}
	default:
		log.Fatalln("Tipo não definido.")
	}
	return
}

func add(fim *[][11]int, ret ...retorno) {
	var enc [11]int
	var pre = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, r := range ret {
		enc[r.idx] = r.val
		pre = remove(pre, Index(pre, r.idx))
	}
	for i := 0; i < 10; i++ {
		if len(pre) == 1 {
			enc[pre[0]] = i
			if !ifExist(*fim, enc) {
				*fim = append(*fim, enc)
			}
		} else {
			for j := 0; j < 10; j++ {
				if len(pre) == 2 {
					enc[pre[0]], enc[pre[1]] = i, j
					if !ifExist(*fim, enc) {
						*fim = append(*fim, enc)
					}
				} else {
					for k := 0; k < 10; k++ {
						if len(pre) == 3 {
							enc[pre[0]], enc[pre[1]], enc[pre[2]] = i, j, k
							if !ifExist(*fim, enc) {
								*fim = append(*fim, enc)
							}
						}
					}
				}
			}
		}
	}
}

func ifExist(ini [][11]int, nova [11]int) bool {
	for _, z := range ini {
		teste := true
		for y := 0; y < 11; y++ {
			if z[y] != nova[y] {
				teste = false
				break
			}
		}
		if teste {
			return true
		}
	}
	return false
}

func ifIn(ini []int, t int) bool {
	for _, z := range ini {
		if z == t {
			return true
		}
	}
	return false
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
func Index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func cpfToString(cpf [11]int) (str string) {
	for _, v := range cpf {
		t := strconv.Itoa(v)
		str += t
	}
	return
}

func testes(num [][11]int) (cpf, nis []string) {
	for _, ints := range num {
		if termo == 0 || termo == 1 {
			if isCPF(ints) {
				cpf = append(cpf, cpfToFormated(ints))
			}
		}
		if termo == 0 || termo == 2{
			if isNIS(ints) {
				nis = append(nis, nisToFormated(ints))
			}
		}
	}
	sort.Strings(cpf)
	sort.Strings(nis)
	return
}

func tabula(val, i int, s string) string {
	size, _ := ts.GetSize()
	switch {
	case val < 10:
		tab = size.Col() / 20
		return fmt.Sprintf("   %d: %s", i+1, s)
	case val < 100:
		tab = size.Col() / 21
		return fmt.Sprintf("   %02d: %s", i+1, s)
	case val < 1000:
		tab = size.Col() / 22
		return fmt.Sprintf("   %03d: %s", i+1, s)
	case val < 10000:
		tab = size.Col() / 23
		return fmt.Sprintf("   %04d: %s", i+1, s)
	case val < 100000:
		tab = size.Col() / 24
		return fmt.Sprintf("   %05d: %s", i+1, s)
	case val < 1000000:
		tab = size.Col() / 25
		return fmt.Sprintf("   %06d: %s", i+1, s)
	case val < 10000000:
		tab = size.Col() / 26
		return fmt.Sprintf("   %07d: %s", i+1, s)
	case val < 100000000:
		tab = size.Col() / 27
		return fmt.Sprintf("   %08d: %s", i+1, s)
	default:
		tab = size.Col() / 20
		return fmt.Sprintf("   %d: %s", i+1, s)
	}
}

func resultado(fim [][11]int) {
	cpf, nis := testes(fim)
	ncpf, nnis := len(cpf), len(nis)
	if ncpf != 0 && nnis != 0 {
		fmt.Printf("\n   * PROVÁVEIS NÚMEROS PARA PIS E CPF *\n")
	}
	if ncpf == 0 && nnis == 0 {
		fmt.Printf("\n   * NENHUM PROVÁVEL NÚMERO CALCULADO PARA PIS OU CPF *\n")
	}
	if ncpf == 0 && nnis != 0 {
		fmt.Printf("\n   * PROVÁVEIS NÚMEROS PARA PIS *\n")
	}
	if ncpf != 0 && nnis == 0 {
		fmt.Printf("\n   * PROVÁVEIS NÚMEROS PARA CPF *\n")
	}
	if ncpf != 0 {
		fmt.Printf("\n   Para números de CPF:\n\n")
		for i, s := range cpf {
			fmt.Printf("%s", tabula(ncpf, i, s))
			if i%tab == tab-1 {
				fmt.Println()
			}
		}
		if (ncpf-1)%tab != tab-1 {
			fmt.Println()
		}
		fmt.Println("   Total de CPFs:", len(cpf))
	}

	if nnis != 0 {
		fmt.Printf("\n   Para números de NIS:\n\n")
		for i, s := range nis {
			fmt.Printf("%s", tabula(nnis, i, s))
			if i%tab == tab-1 {
				fmt.Println()
			}
		}
		if (nnis-1)%tab != tab-1 {
			fmt.Println()
		}
		fmt.Println("   Total de NIS:", len(nis))
	}
	if ncpf != 0 && nnis != 0 {
		fmt.Println("\n   Total:", len(cpf)+len(nis))
	}
}
