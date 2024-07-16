# Gea 2

Es un tablero para jugar juegos de rol online como por ejemplo roll20. Escribo este código tan solo para
familiarizarme con Golang.

# Mis recursos para seguir aprendiendo con tutoriales

- [Tutorial](https://github.com/anaseto/gruid-rltuto/tree/part-1)
- [Tutorial Rust](https://www.reddit.com/r/roguelikedev/comments/101q4pl/rust_roguelike_tutorial_postmortem/)
- [15 steps](http://www.roguebasin.com/index.php?title=How_to_Write_a_Roguelike_in_15_Steps)
- [Tutorial in golang](https://www.fatoldyeti.com/posts/roguelike-tutorial-0/)

# Errores comunes

- Falta un campo en el mapa 01.yaml, en mi caso concreto faltaba el campo "i" de una casilla
- No pinta nada en el lienzo en las coordenadas 0,0 porque deben empezar en 1,1
