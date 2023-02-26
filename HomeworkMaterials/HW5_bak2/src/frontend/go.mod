module dclass.ucscx.edu/frontend

go 1.20

replace dclass.ucscx.edu/gorilla/mux => ../gorilla/mux

replace dclass.ucscx.edu/contract => ../contract

replace dclass.ucscx.edu/lib => ../lib

require (
	dclass.ucscx.edu/contract v0.0.0-00010101000000-000000000000
	dclass.ucscx.edu/gorilla/mux v0.0.0-00010101000000-000000000000
	dclass.ucscx.edu/lib v0.0.0-00010101000000-000000000000
)

require github.com/gorilla/mux v1.8.0 // indirect
