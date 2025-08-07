CURSO DE PROGRAMACION ORIENTADA A OBJETOS
Pilares de programacion en GO   
Abstraccion:   Proceso mental realizado para aislar un elemento del mundo real o fictisio
                Un ejemplo es cuando se subtrae de la realidad caracteristicas de un carro para ponerlas dentro de un codigo 
                tambien se puede abstraer caracteristicas de cosas ficticias

                en OOP es la capacidad de implementar las caracteristicas en codigo en C++/java se usan las clases con sus atributos y metodos

                en Go se usan las estructuras con sus campos de estructuras y metodos, ademas de que los metodos no son propios de las estructuras si no que pueden ser utilizados en distintas estructuras, mapas, slices, strings , etc

Encapsulamiento:  Caracteristica que proove un lenguaje para protejer las propiedades o metodos con los que trabaja el programador
                    ejemplos fuera de go son private, public, protected
                    mientras en go se utilizan los identificadores exportados o no exportados

                    La encapsulacion permite encapsular la complejidad de un proceso

Herencia:       es la capacidad de transmitir caracteristicas(propiedades y metodos) de una clase a otra, por medio del sistema de herencia en 
Composicion     otros lenguajes como c++/java
                En GO no se cuenta con esta caracteristica como tal, se tiene un sistema de tipos con de se plasman caracteristicas y metodos que permiten trabajar de manera similar, este sistema permite usar tipos que pertenescan a otros tipos de manera analoga pero distinta
                ademas de que existen el concepto de interfaces

                desde el diseno del lenguaje se diseno para que no existiera la herencia, esto para evitar la gerarquia de tipos y un sobrediseno temprano del programa 

                Como alternativa se genero el concepto de composicion, a diferencia de la herencia donde la relacion es de uno a uno, es decir una clase hereda solo de uno, mientras que en la composicion es relaciones 1 a muchos, no hay gerarquia de tipos, si no mas bien principios de responsabilidad, ademas de que se pueden agragar dependencias

                la composicion permite la realizacion de test unitarios de manera sensilla, al generar metodos que testien sola mente la interfas que se busquen.

                otra ventaja es de que se pueden tener metodos en diferentes tipos declarados ademas de estructuras, ademas de que se pueden tener interfaces implicitas, esto permite que un tipo pueda satisfacer diferentes interfaces

                Go permite inscruptar tipos dentro de otros tipos similar a las subclases



Polimorfismo:       La caracteristica del lenguaje de mandar para mandar mensajes sintacticamente iguales y que objetos de distintos tipos puedan 
                reaccionar al mensaje, a diferencia de otros lenguajes esto se realiza en GO a travez del uso de interfaces


Tipos Declarados:   Son nuevos tipos creados por el programador los cuales son creados a base de tipos existentes, con la condicion de que estos tipos,
                    sus metodos, deben de estar declarados en el paquete donde se han sido declarados, razon por la cual no es posible agregar metodos a los tipos predeclarados como bool, string, byte, ya que no se han declarado en nuestro paquete, sin embargo se puede crear un nuevo tipo con base en un tipo predeclarado y que agregar los metodos que los tengamos

                   


