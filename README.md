# Terrytalks

Generates Terry Davis phrases based on a Markov chain algorithm. The Markov  
chain is built from transcripts of video logs and recordings of live streams  
that Terry uploaded to youtube.  

To build:  
```$ go generate```  
```$ go install github.com/bitwitch/terrytalks```  

```
The following command line options are available:  

-words        maximum number of words to print  

-prefix       prefix length in words  

Examples: terrytalks 
          terrytalks -words=10
          terrytalks -words=1000 -prefix=4
```

---
I love Terry Davis. He was an honest, intelligent man and was wise in spite  
of his schizophrenia. He had his issues, but I focus on the positive he  
brings to the world.  

Sometimes when he communicated online, he used a markov chain text generator  
presumably seeded with text from the bible. At times he seemed to revere it  
as divine truth, and at [other times](https://www.metafilter.com/119424/An-Operating-System-for-Songs-from-God#4538454) he seemed to see it as silly and arbitrary. 