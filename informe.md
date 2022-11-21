#  Informe Algogram 



### Introducción & Análisis del Problema
Comenzamos analizando los problemas principales que debimos afrontar al principio del trabajo. ¿Qué estructura de datos deberíamos usar?, ¿Qué TDA's deberíamos implementar?, ¿Qué criterio vamos a tomar con respecto a los ingresos pasados por parámetro? Ahora en breve explicaremos detalladamente todo…

Teniendo en cuenta la consigna, decidimos crear un diagrama de flujo para poder definir una estructura básica de cómo afrontar la implementación de la red social. 

Para aclarar, no definimos nuestro TP en base a este diagrama, solo fue hecho para poder tener una idea de como realizar el trabajo práctico.

A continuación dividimos el programa en un conjunto de subproblemas:
<hr>
1- Necesitamos que las operaciones de “login” y “logout” sean O(1) en complejidad temporal por lo que debemos usar algún tipo de estructura de datos, ya que no nos sirve leer el archivo de usuarios cada vez que dichas operaciones se lleven a cabo.




2 - Necesitamos afrontar el tema de los comandos “Publicar Post” y “Ver próximo post en el feed”. Para esto, decidimos implementar 2 TDA’s distintos, uno para los usuarios y otro para los posts. Así podemos crear múltiples posts y asignarlos a distintos usuarios.

3- Para llevar a cabo operaciones de “likear” y “mostrar likes” necesitamos guardar esa información. Para esto vamos a usar una estructura de datos que estará en memoria hasta el momento de terminación del programa (ya que no se pide que se almacenen los datos) y que nos ofrezca la complejidad temporal especificada en la consigna.

### Solución propuesta

1- Nuestra solución a esta problemática fue; Leer una sola vez el archivo de usuarios dado por parámetro al ejecutar el programa, y por cada usuario dentro del archivo instanciar un nuevo Usuario TDA (cuyo comportamiento está explicado más adelante en el informe) y guardar estos en un hashmap por nombre de usuario de manera que puedan ser accedidos en tiempo constante al momento de realizar un “login”, en contraste a este comando al realizar un “logout” no usaremos el hash ya que es tan simple como reinicializar una variable.

2- Con respecto a la implementación de los TDA’s de usuarios y posts. Nuestra solución consiste en lo siguiente:
Todos los usuarios tienen su propio heap que sirve para determinar qué posts aparecerán primero en su feed, comparando los posts con respecto a el “nivel de afinidad”*  de sus usuarios dueños. Cuando un usuario decide crear un post, este post será agregado al feed de todos los demás usuarios (para esto se itera el hashmap de usuarios).
La complejidad de la acción será O(u log(p)) ya que por cada usuario, debemos encolar el post a su propio heap.

*Nivel de Afinidad: 
Es la “distancia” que tiene un usuario con respecto a otro en el archivo de usuarios.
Cuanto menor sea la distancia, mayor será su afinidad.
Ejemplo:	 
[0] usuarioA 
[1] usuarioB 
[2] usuarioC 

La “distancia” del usuarioC al usuarioA es 2, mientras tanto que la “distancia” de usuarioC a usuarioB será 1, entonces el usuarioC tiene mayor afinidad con el usuarioB en comparación con el usuario A.




3- Al momento que un usuario likea un post debemos guardar en ese post la información de quienes dieron like ordenarlos alfabéticamente con una complejidad temporal de O(log Up) siendo Up la cantidad de personas que dieron like, por lo que decidimos utilizar un árbol binario de búsqueda ya que mantendrá todo ordenado y nos permite además recorrerlo completo para mostrar los likes en tiempo lineal (recorriendo todos los likers). Para obtener rápidamente el post (y no alentar las operaciones), utilizamos un segundo hashmap, en este caso por id del post en cuestión.


#### TDA Usuario

Se definió un Usuario de tal manera que este tenga la información de su nombre, feed y su “nivel de afinidad con otros usuarios”, y se le dio un comportamiento tal que este pueda ver su feed de posts, así como guardarlos en la misma, y también pueda ser capaz de dar su nombre y nivel de afinidad con fines de comparar al nivel de afinidad de otro usuario de algogram.


#### TDA Post

Al crear un post, se guarda la información de la descripción del post, de su dueño, de los likes y el ID del post. Decidimos guardar información del dueño al post para cuando se requiera agregar un post al feed podemos comparar el “nivel de afinidad” de un usuario con respecto a otro.

