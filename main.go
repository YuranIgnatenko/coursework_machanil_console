package main

import (
	// "os"
	// "math"
	"fmt"
)

var (
	K_700_NKR_U = 3200
	K_700_NKR_N = 11970
	K_700_NTR_N = 5120
	K_700_NTO_3_N = 2560
	K_700_NTO_2_N = 640
	K_700_NTO_1_N = 160
	K_700_NCTO = 2

	K_701_NKR_U = 3500
	K_701_NKR_N = 19220
	K_701_NTR_N = 6200
	K_701_NTO_3_N = 3120
	K_701_NTO_2_N = 780
	K_701_NTO_1_N = 195
	K_701_NCTO = 2

	DT_NKR_U = 1150
	DT_NKR_N = 5800
	DT_NTR_N = 2480
	DT_NTO_3_N = 1240
	DT_NTO_2_N = 310
	DT_NTO_1_N = 77
	DT_NCTO = 2

	MTZ2_NKR_U = 1300
	MTZ2_NKR_N = 4480
	MTZ2_NTR_N = 1680
	MTZ2_NTO_3_N = 840
	MTZ2_NTO_2_N = 210
	MTZ2_NTO_1_N = 52
	MTZ2_NCTO = 2

	MZT_NKR_U = 1200
	MZT_NKR_N = 4480
	MZT_NTR_N = 1680
	MZT_NTO_3_N = 840
	MZT_NTO_2_N = 210
	MZT_NTO_1_N = 52
	MZT_NCTO = 2

	T40_NKR_U = 700
	T40_NKR_N = 2630
	T40_NTR_N = 1200
	T40_NTO_3_N = 600
	T40_NTO_2_N = 150
	T40_NTO_1_N = 37
	T40_NCTO = 2

	T25_NKR_U = 350
	T25_NKR_N = 1940
	T25_NTR_N = 740
	T25_NTO_3_N = 370
	T25_NTO_2_N = 92
	T25_NTO_1_N = 23
	T25_NCTO = 2

	T16_NKR_U = 200
	T16_NKR_N = 1320
	T16_NTR_N = 510
	T16_NTO_3_N = 255
	T16_NTO_2_N = 64
	T16_NTO_1_N = 16
	T16_NCTO = 2

	NTR_BDT = 0.78
	NCTO_BDT = 2

	NTR_SZ = 0.78
	NCTO_SZ = 2

	NTR_KKU = 0.8
	NCTO_KKU = 2

	CK_NKR = 0.15
	CK_NTR = 0.6
	CK_NTO = 220
	CK_NTO_2 = 240
	CK_NTO_1 = 60
	CK_NCTO = 1

	UAZ_NKR_U = 35000
	UAZ_NKR_N = 140000
	UAZ_NTO_2_N = 12800
	UAZ_NTO_1_N = 3200
	UAZ_NCTO = 2

	KAMAZ_NKR_U = 40000
	KAMAZ_NKR_N = 250000
	KAMAZ_NTO_2_N = 7200
	KAMAZ_NTO_1_N = 2400
	KAMAZ_NCTO = 2

	GAZ_NKR_U = 36000
	GAZ_NKR_N = 160000
	GAZ_NTO_2_N = 10000
	GAZ_NTO_1_N = 2500
	GAZ_NCTO = 2

	ZIL_NKR_U = 11000
	ZIL_NKR_N = 230000
	ZIL_NTO_2_N = 10000
	ZIL_NTO_1_N = 2500
	ZIL_NCTO = 2

	K_T_KR = 410
	K_T_TR = 297
	K_T_TO_3 = 43.2
	K_T_TO_2 = 10.6
	K_T_TO_1 = 2.5
	K_T_CTO = 29.3

	K1_T_KR = 451
	K1_T_TR = 297
	K1_T_TO_3 = 25.2
	K1_T_TO_2 = 11.6
	K1_T_TO_1 = 2.2
	K1_T_CTO = 182.3

	T_KR_DT = 229
	T_TR_DT = 268
	T_TO_3_DT = 21.4
	T_TO_2_DT = 6.4
	T_TO_1_DT = 2.7
	T_CTO_DT = 17.1

	T_KR_MTZ2 = 193
	T_TR_MTZ2 = 163
	T_TO_3_MTZ2 = 19.8
	T_TO_2_MTZ2 = 6.9
	T_TO_1_MTZ2 = 2.7
	T_CTO_MTZ2 = 3.5

	T_KR_MZT = 229
	T_TR_MZT = 268
	T_TO_3_MZT = 19.8
	T_TO_2_MZT = 6.9
	T_TO_1_MZT = 2.7
	T_CTO_MZT = 3.5

	T_KR_T40 = 156
	T_TR_T40 = 127
	T_TO_3_T40 = 18
	T_TO_2_T40 = 6.8
	T_TO_1_T40 = 2
	T_CTO_T40 = 19.8

	T_KR_T25 = 132
	T_TR_T25 = 80
	T_TO_3_T25 = 7.7
	T_TO_2_T25 = 2.7
	T_TO_1_T25 = 0.9
	T_CTO_T25 = 1.8


	T_KR_T16 = 114
	T_TR_T16 = 80
	T_TO_3_T16 = 7.7
	T_TO_2_T16 = 2.7
	T_TO_1_T16 = 0.9
	T_CTO_T16 = 1.8

	T_TR_KKU = 69
	T_CTO_KKU = 3.4

	T_TR_PLN = 17
	T_CTO_PLN = 3.4

	T_TR_BDT = 29
	T_CTO_BDT = 3.4

	T_TR_SZ = 29
	T_CTO_SZ = 3.4

	CK_T_KR = 330
	CK_T_TR = 150
	CK_T_TO_2 = 6.6
	CK_T_TO_1 = 5.1
	CK_T_CTO = 15

	UAZ_T_KR = 280
	UAZ_T_TO_2 = 11.1
	UAZ_T_TO_1 = 2.2
	UAZ_T_CTO = 0.25 * 7.9
	UAZ_T_TR_U1 = 35000
	UAZ_T_TR_D = 1000
	UAZ_T_TR_Z = 10.3

	KAMAZ_T_KR = 380
	KAMAZ_T_TO_2 = 21.5
	KAMAZ_T_TO_1 = 4.4
	KAMAZ_T_CTO = 0.25 * 10.2
	KAMAZ_T_TR_U1 = 40000
	KAMAZ_T_TR_D = 1000
	KAMAZ_T_TR_Z = 10.5

	GAZ_T_KR = 249
	GAZ_T_TO_2 = 11.8
	GAZ_T_TO_1 = 2.9
	GAZ_T_CTO = 0.25 * 4.6
	GAZ_T_TR_U1 = 36000
	GAZ_T_TR_D = 1000
	GAZ_T_TR_Z = 6

	ZIL_T_KR = 302
	ZIL_T_TO_2 = 14
	ZIL_T_TO_1 = 3.5
	ZIL_T_CTO = 0.25 * 5
	ZIL_T_TR_U1 = 44000
	ZIL_T_TR_D = 1000
	ZIL_T_TR_Z = 6.3
)

// var int array_name_count["K-700"][1]
// var int K_701
// var int DT
// var int MTZ2
// var int MTZ
// var int T40
// var int T25
// var int T16
// var int PLN
// var int BDT
// var int SZ

// var array_var_count = []*int{
// 	*array_name_count["K-700"][1],
// 	*K_701,
// 	*DT,
// 	*MTZ2,
// 	*MTZ,
// 	*T40,
// 	*T25,
// 	*T16,
// 	*PLN,
// 	*BDT,
// 	*SZ,
// }

// все ключи мапы на русском языке
var array_name_count = map[string][]int{
	"К_700":       {0,2},
	"К_701":       {1,2},
	"ДТ_75_МВ":    {2,2},
	"МТЗ_82":      {3,2},
	"МТЗ_80":      {4,2},
	"Т_40":        {5,2},
	"Т_25":        {6,2},
	"Т_16":        {7,2},
	"ПЛН_4_35":    {8,2},
	"БДТ_3":       {9,2},
	"СЗ_3_6":      {10,2},
	"ККУ_2А":      {11,2},
	"СК_5":        {12,2},
	"УАЗ_469":     {13,2},
	"КАМАЗ_5320":  {14,2},
	"ГАЗ_53":      {15,2},
	"ЗИЛ_130":     {16,2},
}

var array_result_formula_N = make([]int, 0)

func View(data interface{}){
	fmt.Println(data)
}

func ScanCountMachine(){	
	var user_count int
	for i:=0; i<len(array_name_count); i++ {
		for k,v := range array_name_count{
			if v[0] == i{
				fmt.Print("ENTER COUNT ", k ,": ")
				dt := make([]int,0)
				dt = append(dt, i)
				fmt.Scan(&user_count)
				dt = append(dt, user_count)
				array_name_count[k] = dt
			}
		}
	}
}

func int_r (x float64) int {
	// Round возвращает ближайшее целочисленное значение.
	if x >= float64(int(x))+0.5{
		return int(x+1)
	}else{
		return int(x)
	}
}

// array_name_count["K-700"][1]

func itf(x int) float64{
	return float64(x)
}

func Counting_N (ARRAY_KEY_COUNT map[string][]int,KEY string, MACHINE_NKR_U int, MACHINE_NKR_N int ,MACHINE_NTR_N int , MACHINE_NTO_3_N int , MACHINE_NTO_2_N int , MACHINE_NTO_1_N int , MACHINE_NCTO int ) (int,int,int,int,int,int){
	COUNT_M := ARRAY_KEY_COUNT[KEY][1]

	NKR_MACHINE := itf(MACHINE_NKR_U) * itf(COUNT_M) / itf(MACHINE_NKR_N)
	NKR_MACHINE_d := int_r(NKR_MACHINE) 
	NTR_MACHINE := itf(MACHINE_NKR_U) * itf(COUNT_M) / itf(MACHINE_NTR_N) - itf(NKR_MACHINE_d)
	NTR_MACHINE_d := int_r(NTR_MACHINE) 
	NTO_3_MACHINE := itf(MACHINE_NKR_U) * itf(COUNT_M) / itf(MACHINE_NTO_3_N) - itf(NKR_MACHINE_d) - itf(NTR_MACHINE_d)
	NTO_3_MACHINE_d := int_r(NTO_3_MACHINE) 
	NTO_2_MACHINE := itf(MACHINE_NKR_U) * itf(COUNT_M) / itf(MACHINE_NTO_2_N) - itf(NKR_MACHINE_d) - itf(NTR_MACHINE_d) - itf(NTO_3_MACHINE_d)
	NTO_2_MACHINE_d := int_r(NTO_2_MACHINE) 
	NTO_1_MACHINE := itf(MACHINE_NKR_U) * itf(COUNT_M) / itf(MACHINE_NTO_1_N) - itf(NKR_MACHINE_d) - itf(NTR_MACHINE_d) - itf(NTO_3_MACHINE_d) - itf(NTO_2_MACHINE_d)
	NTO_1_MACHINE_d := int_r(NTO_1_MACHINE) 
	MACHINE_N_CTO := MACHINE_NCTO * COUNT_M

	// print("Nто-3"+" = "+str(K_701_NKR_U)+" * "+str(K_701)+" / "+str(K_701_NTO_3_N)+" - "+str(NKR_K_701)+" - "+str(NTR_K_701)+" = "+str(NTO_3_K_701))
	// print ("ПРИНИМАЮ "+ str(NTO_3_K_701))
	// print("Nто-2"+" = "+str(K_701_NKR_U)+" * "+str(K_701)+" / "+str(K_701_NTO_2_N)+" - "+str(NKR_K_701)+" - "+str(NTR_K_701)+" - "+str(NTO_3_K_701)+" = "+str(NTO_2_K_701))
	// print ("ПРИНИМАЮ "+ str(NTO_2_K_701))
	// print("Nто-1"+" = "+str(K_701_NKR_U)+" * "+str(K_701)+" / "+str(K_701_NTO_1_N)+" - "+str(NKR_K_701)+" - "+str(NTR_K_701)+" - "+str(NTO_3_K_701)+" - "+str(NTO_2_K_701)+" = "+str(NTO_1_K_701))
	// print ("ПРИНИМАЮ "+ str(NTO_1_K_701))
	// print("Nсто"+" = "+str(K_701_NCTO)+" * "+str(K_701)+" = "+str(K_701_N_CTO))
	
	fmt.Println("\n",KEY," :")

	fmt.Printf("NKR = %f * %f / %f = %f\nПринимаю: %d\n",
	MACHINE_NKR_U,COUNT_M,MACHINE_NKR_N, NKR_MACHINE, NKR_MACHINE_d)

	fmt.Printf("NTR = %f * %f / %f - %f = %f\nПринимаю: %d\n",
	MACHINE_NKR_U,COUNT_M,MACHINE_NTR_N, NKR_MACHINE_d, NTR_MACHINE, NTR_MACHINE_d)

	fmt.Printf("NT0-3 = %f * %f / %f - %f - %f = %f\nПринимаю: %d\n",
	MACHINE_NKR_U,COUNT_M,MACHINE_NTO_3_N, NKR_MACHINE_d, NTR_MACHINE_d, NTO_3_MACHINE, NTO_3_MACHINE_d)

	fmt.Printf("NT0-2 = %f * %f / %f - %f - %f - %f = %f\nПринимаю: %d\n",
	MACHINE_NKR_U,COUNT_M,MACHINE_NTO_2_N, NKR_MACHINE_d, NTR_MACHINE_d, NTO_3_MACHINE_d, NTO_2_MACHINE, NTO_2_MACHINE_d)

	fmt.Printf("NT0-1 = %f * %f / %f - %f - %f - %f - %f = %f\nПринимаю: %d\n",
	MACHINE_NKR_U,COUNT_M,MACHINE_NTO_2_N, NKR_MACHINE_d, NTR_MACHINE_d, NTO_3_MACHINE_d, NTO_2_MACHINE_d, NTO_1_MACHINE, NTO_1_MACHINE_d)

	fmt.Printf("NCTO = %d * %d = %d\n\n",
	MACHINE_NCTO, COUNT_M, MACHINE_N_CTO)

	return NKR_MACHINE_d, NTR_MACHINE_d, NTO_3_MACHINE_d, NTO_2_MACHINE_d, NTO_1_MACHINE_d, MACHINE_N_CTO
}



func main() {
	View("STARTED PROGRAMM")
	// ScanCountMachine()
	// fmt.Println(array_name_count)
	// fmt.Println(int_r(1.1))
	// fmt.Println(int_r(1.5))
	// fmt.Println(int_r(1.9))
	// fmt.Println(int_r(3.2))
	// N_K_700ff()
	Counting_N(array_name_count, "К_700" ,K_700_NKR_U, K_700_NKR_N, K_700_NTR_N, K_700_NTO_3_N, K_700_NTO_2_N, K_700_NTO_1_N, K_700_NCTO)
	Counting_N(array_name_count, "К_701" ,K_701_NKR_U, K_701_NKR_N, K_701_NTR_N, K_701_NTO_3_N, K_701_NTO_2_N, K_701_NTO_1_N, K_701_NCTO)
	Counting_N(array_name_count, "ДТ_75_МВ" ,DT_NKR_U, DT_NKR_N, DT_NTR_N, DT_NTO_3_N, DT_NTO_2_N, DT_NTO_1_N, DT_NCTO)
	Counting_N(array_name_count, "МТЗ_82" ,MTZ2_NKR_U, MTZ2_NKR_N, MTZ2_NTR_N, MTZ2_NTO_3_N, MTZ2_NTO_2_N, MTZ2_NTO_1_N, MTZ2_NCTO)
	Counting_N(array_name_count, "МТЗ_80" ,MZT_NKR_U, MZT_NKR_N, MZT_NTR_N, MZT_NTO_3_N, MZT_NTO_2_N, MZT_NTO_1_N, MZT_NCTO)

	Counting_N(array_name_count, "Т_40" ,T40_NKR_U, T40_NKR_N, T40_NTR_N, T40_NTO_3_N, T40_NTO_2_N, T40_NTO_1_N, T40_NCTO)
	Counting_N(array_name_count, "Т_25" ,T25_NKR_U, T25_NKR_N, T25_NTR_N, T25_NTO_3_N, T25_NTO_2_N, T25_NTO_1_N, T25_NCTO)
	Counting_N(array_name_count, "Т_16" ,T16_NKR_U, T16_NKR_N, T16_NTR_N, T16_NTO_3_N, T16_NTO_2_N, T16_NTO_1_N, T16_NCTO)

}