# Graphs
Graphs in computer science is a concept that is commonly found in mathematics with many uses in computer science. 
Graphs come in many different flavors, many of which have found uses in computer programs. 

Here are a few just as a reference:
* Simple graph
* Undirected or directed graphs
* Cyclic or acyclic graphs
* labeled graphs
* Weighted graphs
* Infinite graphs

Most graphs follow a variation of these rules: 

* A graph is made up of a set of Vertices and a set of Edges.
* The Verticies are of some underlying type, and the set may be finite or infinite.
* Their are two elements of an Edge that make up a pair from the set of Vertices.
* Graphs are often depicted visually, by drawing the elements of the Vertices set as boxes or circles. These can be refered to as nodes.  
The elements of the edge set are depicted as lines or arcs connecting the nodes.
There is an arc between v1 and v2 if (v1,v2) is an element of the Edge set.

Lets take a look at the *Adjaceny* keyword:  
If **(u,v)** is in the edge set we say **u** is *adjacent* to **v** (which we sometimes write as u ~ v).

<img title="simple graph" alt="Dead =[" src="https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fcourses.cs.vt.edu%2Fcsonline%2FDataStructures%2FLessons%2FGraphs%2Fgraph.gif&f=1&nofb=1">

Take a look at this example above ^

This graph has the following characteristics:
* The underlying set for the Vertices are the characters.
* The Vertices set == **{'A', 'B', 'C', 'D', 'E', 'F'}**
* The Edge set are the following unordered pairs == { **('A', 'B'), ('A', 'C'), ('B', 'C'), ('A', 'D'), ('B', 'E'), ('D', 'E'), ('D', 'F')** }
