package main

import("fmt"
       "strings"
	   "bufio"
	   "os"
	   "log")

func MKDISK(x string, y string, z string, w string){
	//fmt.Println("Este es: "+x+ " junto con: "+y+ " seguido por: "+z+" finalizando para: "+w  )

	var igual string = "=";

	var espacio string = " ";

	var cont int = 0;

	var cont1 int = 0;

	var mknum [10]string;

	var mknum1 [3]string;

	mknum1[0] = "null";
	mknum1[1] = "null";
	mknum1[2] = "null";

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

	split1 := strings.Split(x, igual)

	for _, elec1 := range split1 {
		fmt.Println(elec1)
		mknum1[cont1] = elec1;
		cont1 = cont1 + 1; 
	}

	cont1 = 0;

	mknum1[0] = strings.ToLower(mknum1[0])
	//fmt.Println(mknum[0])
}

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

func PAUSE(){
	file, err := os.Open( "/home/manuel/Escritorio/parte-crear-discos.script")
	if err != nil {
	  log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	  fmt.Println(scanner.Text())
	  fmt.Println("HOLA SI PASO-------------------------")
	}
}

func EXEC(x string){

	var igual string = "=";

	var espacio string = " ";

	var cont int = 0;

	var cont1 int = 0;

	var mknum [10]string;

	var mknum1 [3]string;

	mknum1[0] = "null";
	mknum1[1] = "null";
	mknum1[2] = "null";

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

	split1 := strings.Split(x, igual)

	for _, elec1 := range split1 {
		fmt.Println(elec1)
		mknum1[cont1] = elec1;
		cont1 = cont1 + 1; 
	}

	cont1 = 0;

	mknum1[0] = strings.ToLower(mknum1[0])
	//fmt.Println(mknum[0])

	if mknum1[0] == "-path"{
		fmt.Println("Parametro de comando correcto, iniciando accion...")

		file, err := os.Open(mknum1[1])
		if err != nil {
		  log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)


		for scanner.Scan() {
		  fmt.Println(scanner.Text())
		  //fmt.Println("Hola si paso")

		  split := strings.Split(scanner.Text(), espacio)

		  for _, elec := range split {
			  fmt.Println(elec)
			  mknum[cont] = elec;
			  cont = cont + 1; 
		  }
		  
  
		  mknum[0] = strings.ToLower(mknum[0])
  
		  switch mknum[0]{
  
		  case "mkdisk":
			  MKDISK(mknum[1], mknum[2], mknum[3], mknum[4])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  mknum[4] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "rmdisk":
			  RMDISK(mknum[1])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
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
			  fmt.Println("\n");
  
		  case "mount":
			  MOUNT(mknum[1], mknum[2])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "mkfs":
			  MKFS(mknum[1], mknum[2])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "login":
			  LOGIN(mknum[1], mknum[2], mknum[3])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "logout":
			  LOGOUT()
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			   mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "mkgrp":
			  MKGRP(mknum[1])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "rmgrp":
			  RMGRP(mknum[1])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "mkuser":
			  MKUSER(mknum[1], mknum[2], mknum[3])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "rmusr":
			  RMUSR(mknum[1])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
	  
		  case "mkfile":
			  MKFILE(mknum[1], mknum[2], mknum[3], mknum[4])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  mknum[4] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "mkdir":
			  MKDIR(mknum[1], mknum[2])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "pause":
			  PAUSE();
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "exec":
			  EXEC(mknum[1])
  
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  cont = 0;
			  fmt.Println("\n");
  
		  case "rep":
			  REP(mknum[1], mknum[2], mknum[3], mknum[4])
			  mknum[0] = "null";
			  mknum[1] = "null";
			  mknum[2] = "null";
			  mknum[3] = "null";
			  mknum[4] = "null";
			  cont = 0;
			  fmt.Println("\n");
	  
		  case "exit":
			  fmt.Println("Saliendo del programa")
			  cont = 0;
			  fmt.Println("\n");
		  default:
			  fmt.Println("Comando no identificado, Intente otra vez...")
			  cont = 0;
			  fmt.Println("\n");
  
		  }
		}



	}else{
		fmt.Println("ERROR-Parametro de comando incorrecto")
	}

}

func REP(x string, y string, z string, w string){

}

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


		split := strings.Split(eleccion, espacio)

		for _, elec := range split {
			fmt.Println(elec)
			mknum[cont] = elec;
			cont = cont + 1; 
		}
		

		mknum[0] = strings.ToLower(mknum[0])

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


		if mknum[0] == "exit"{
			break
		}

	}






}