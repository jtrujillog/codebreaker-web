Hola Juan Felipe # codebreaker-web
Para instalar paquetes en go se hace con #go get github.com/gin-gonic/gin
go get funciona como npm install para node js
Si saca este error: $GOPATH not set
Debemos configurar la variable de ambiente go path

Editar el archivo /etc/bash.bashrc y al final poner las líneas
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

Cerrar la consola y volverla a abrir
Crear la carpeta go en $HOME
$ cd $GOPATH Si funciona entonces la variable quedó bien creada

Ejecutamos nuevamente #go get github.com/gin-gonic/gin
Luego entramos a la carpeta del api y ejecutamos #go build
Si saca errores de código se corrigen y se vuelve a ejecutar
luego $ls para ver lo que creo, debe crear un archivo ejecutable y lo corremos con $./<nombre archivo creado con el build> esto nos corre el servidor

ingresamos a la página  http://localhost:8081/codebreaker/setup/<number secreto>
luego http://localhost:8081/codebreaker/guess/<number a comparar>
