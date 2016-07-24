# swatchout
A golang CLI for extracting a color palette from an image

This program will read an image from disk, and then produce a given number of json-formatted RGBA values to use for something pretty.

To use this tool, install Go, build a binary using `go build`, and place this somewhere in your `PATH`.

You can invoke it with the following options:

```
swatchout -numcolors={integer} {/path/to/image}
```

For an example, a simple demonstration with [Edmund Tarbell's excellent painting from 1907, 'Preparing for the Matinee'](https://upload.wikimedia.org/wikipedia/commons/thumb/7/74/Edmund_Tarbell_matinee.jpg/476px-Edmund_Tarbell_matinee.jpg).

![Edmund Tarbell, 1907 - "Preparing for the Matinee"](https://upload.wikimedia.org/wikipedia/commons/thumb/7/74/Edmund_Tarbell_matinee.jpg/476px-Edmund_Tarbell_matinee.jpg)

Using swatchout on this image with `numcolors` set to `15` produces the following array of RGBA values:

```
[{"r":89,"g":84,"b":51,"a":1},{"r":177,"g":168,"b":146,"a":1},{"r":166,"g":156,"b":133,"a":1},{"r":207,"g":196,"b":173,"a":1},{"r":235,"g":230,"b":212,"a":1},{"r":110,"g":103,"b":69,"a":1},{"r":140,"g":125,"b":92,"a":1},{"r":156,"g":143,"b":117,"a":1},{"r":28,"g":36,"b":18,"a":1},{"r":56,"g":56,"b":32,"a":1}]
```

Another example, using [Titian's famous painting 'Bacchus and Ariadne'](https://upload.wikimedia.org/wikipedia/commons/thumb/b/bb/Titian_-_Bacchus_and_Ariadne_-_Google_Art_Project.jpg/663px-Titian_-_Bacchus_and_Ariadne_-_Google_Art_Project.jpg).

![Titian, 1520-23, "Bacchus and Ariadne"](https://upload.wikimedia.org/wikipedia/commons/thumb/b/bb/Titian_-_Bacchus_and_Ariadne_-_Google_Art_Project.jpg/663px-Titian_-_Bacchus_and_Ariadne_-_Google_Art_Project.jpg)

Using swatchout on this image with `numcolors` set to `20` produces the following array of RGBA values:

```
[{"r":228,"g":221,"b":184,"a":1},{"r":91,"g":66,"b":39,"a":1},{"r":139,"g":160,"b":178,"a":1},{"r":99,"g":110,"b":114,"a":1},{"r":40,"g":51,"b":69,"a":1},{"r":173,"g":143,"b":104,"a":1},{"r":122,"g":140,"b":154,"a":1},{"r":34,"g":31,"b":17,"a":1},{"r":60,"g":80,"b":118,"a":1},{"r":69,"g":68,"b":49,"a":1},{"r":208,"g":173,"b":135,"a":1},{"r":50,"g":42,"b":24,"a":1},{"r":87,"g":86,"b":68,"a":1},{"r":117,"g":101,"b":72,"a":1},{"r":150,"g":112,"b":78,"a":1},{"r":95,"g":121,"b":166,"a":1},{"r":179,"g":188,"b":177,"a":1},{"r":66,"g":52,"b":30,"a":1},{"r":71,"g":96,"b":143,"a":1},{"r":120,"g":81,"b":50,"a":1}]
```

See [this CodePen](http://codepen.io/timothyrourke/pen/dXZNBR) and [this CodePen](http://codepen.io/timothyrourke/pen/ZOoLgY) for a more visual example.
