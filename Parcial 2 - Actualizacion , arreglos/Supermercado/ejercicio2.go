Accion Ej2 es
Ambiente

	movimientosDiarios = registro
		clave = registro
			productoId: N(5)
		fr
		clienteId: N(5)
		tipo: "C","D" //compra devolucion
		cantidad: N(4)
		precioUnitario: N(10)
		total: N(10)
		tipoEnvio: "ENVIO","RETIRO"
	fr

	archMov: archivo de movimientosDiarios ordenado por clave
	mov: movimientosDiarios

	productos = registro
		clave = registro
			productoId: nombre(5)
		fr
		nombre: AN(40)
		descripcion: AN(80)
		rubro: "LIMPIEZA", "CARNICERIA", "VERDULERIA", "BAZAR", "PANADERIA"
	fr

	archProd: archivo de productos indexado por clave
	prod: productos

	a: arreglo de (1...6,1..3) de entero // i = rubro , j= envio
	i,j,rubro,forma,rubroMayor,envioMayor: entero

	Funcion tipo(x:alfanumerico): entero es
		Segun x hacer
			"LIMPIEZA": tipo:= 1
			"CARNICERIA": tipo:= 2
			"VERDULERIA": tipo:= 3
			"BAZAR": tipo:= 4
			"PANADERIA": tipo:= 5
		fs
	fp

	Funcion forma(x:alfanumerico): entero es
		Segun x hacer
			"C": forma:= 1
			"D": forma:= 2
		fs
	fp

Algoritmo
	AbrirE/(archMov) ; AbrirE/(archProd)
	Para i=1 hasta 6 hacer
		Para j=1 hasta 3 hacer
			A(i,j):= 0
		fp
	fp

	Mientras NFDA(mov) hacer
		paq.productoId:= mov.productoId
		tipo(paq.rubro)
		forma(mov.tipo)
		Leer(archProd,prod)
		i:= rubro
		j:= forma
		Si EXISTE entonces
			A(i,j):= A(i,j) + mov.total
			A(i,3):= A(i,3) + mov.total
			A(6,j):= A(6,j) + mov.total
			A(6,3):= A(6,3) + mov.total
		fs

		Si forma = 2 entonces
			devolucion:= devolucion + 1
		fs
	fm

	mayor:= 0

	Para i=1 hasta 5 hacer
		Para j=1 hasta 2 hacer
			Si A(i,j) > 200000 entonces
				Esc("Para el rubro ", i, "de envio",j,"se tuvo un monto de" a
			i,j)
			fs

			Si A(i,j) > mayor entonces
				rubroMayor:= i
				envioMayor:= j
			fs

		fp
	fp



	Esc("Para el rubro ", rubroMayor, "de envio",envioMayor,"se tuvo el mayor monto")

	Para i=1 hasta 5 hacer
		Esc("Para el rubro ", i ,"Se tuvo ", A(i;3)," devoluciones")
	Fp




fa