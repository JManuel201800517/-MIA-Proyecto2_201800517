package main

import("fmt"
       "strings"
	   "bufio"
	   "os")

func MKDISK(x string, y string, z string, w string){}

func RMDISK(x string){}

func FDISK(x string, y string, z string, w string, v string, u string){}

func MOUNT(x string, y string){}

func MKFS(x string, y string){}

func LOGIN(x string, y string, z string){}

func LOGOUT(){}

func MKGRP(x string){}

func RMGRP(x string){}

func MKUSER(x string, y string, z string){}

func RMUSR(x string){}

func MKFILE(x string, y string, z string, w string){}

func MKDIR(x string, y string){}

func PAUSE(){}

func EXEC(x string){}

func REP(x string, y string, z string, w string){}

func main(){
	fmt.Println("Hello, Word, que hubo cuates")

	//var cmd string 
	//var xx string

	var espacio string = " ";

	var cont int = 0;

	var mknum [10]string;

	mknum[0] = "null";
	mknum[1] = "null";
	mknum[2] = "null";
	mknum[3] = "null";
	mknum[4] = "null";
	mknum[5] = "null";
	mknum[6] = "null";
	mknum[7] = "null";
	mknum[8] = "null";
	mknum[9] = "null";
	
	fmt.Println("\n");

	for{
		fmt.Println("Escribir comando adecuado:")
		//fmt.Scanln(&cmd)

		//fmt.Println(cmd)

		reader := bufio.NewReader(os.Stdin)

	    entrada, _ := reader.ReadString('\n') // Leer hasta el separador de salto de línea
	    eleccion := strings.TrimRight(entrada, "\r\n")

		//fmt.Println(entrada)
		fmt.Println(eleccion)

		eleccion = strings.ToLower(eleccion)

		split := strings.Split(eleccion, espacio)

		for _, elec := range split {
			fmt.Println(elec)
			mknum[cont] = elec;
			cont = cont + 1; 
		}

		switch mknum[0]{

		case "mkdisk":
			MKDISK(mknum[1], mknum[2], mknum[3], mknum[4])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        mknum[4] = "null";
	        cont = 0;

		case "rmdisk":
			RMDISK(mknum[1])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "fdisk":
			FDISK(mknum[1], mknum[2], mknum[3], mknum[4], mknum[5], mknum[6])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
         	mknum[4] = "null";
        	mknum[5] = "null";
	        mknum[6] = "null";
	        cont = 0;

		case "mount":
			MOUNT(mknum[1], mknum[2])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "mkfs":
			MKFS(mknum[1], mknum[2])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "login":
			LOGIN(mknum[1], mknum[2], mknum[3])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "logout":
			LOGOUT()

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
         	mknum[3] = "null";
        	cont = 0;

		case "mkgrp":
			MKGRP(mknum[1])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "rmgrp":
			RMGRP(mknum[1])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "mkuser":
			MKUSER(mknum[1], mknum[2], mknum[3])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "rmusr":
			RMUSR(mknum[1])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;
	
		case "mkfile":
			MKFILE(mknum[1], mknum[2], mknum[3], mknum[4])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        mknum[4] = "null";
	        cont = 0;

		case "mkdir":
			MKDIR(mknum[1], mknum[2])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "pause":
			PAUSE();

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "exec":
			EXEC(mknum[1])

			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        cont = 0;

		case "rep":
			REP(mknum[1], mknum[2], mknum[3], mknum[4])
			mknum[0] = "null";
	        mknum[1] = "null";
	        mknum[2] = "null";
	        mknum[3] = "null";
	        mknum[4] = "null";
	        cont = 0;
	
		case "exit":
			fmt.Println("Saliendo del programa")
			cont = 0;
		default:
			fmt.Println("Comando no identificado, Intente otra vez...")
			cont = 0;

		}


		if eleccion == "exit"{
			break
		}

	}






}