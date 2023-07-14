# autocluster
Hierarchical clustuering of any files. Pure go.


## General idea

* access files, detect if file is already zipped, and if so, unzip it
* function to measure distance between file (see paper), based on unzipped file content
* generate distance matrix (int x int -> float 64)
* define a distance interface (input 2 ints, output a float 64)

* given a distance matrix, define a linkage interface between two clusters (input : 2 clusters, output : a float64 distance)
* define a cluster type, that contains :
  * cluster id,
  * slice of objects inds in cluster
  * left, right sub cluster
  * linkage distance between left and right subclusters,
  * potentially, some data to facilitate cluster distance computation (recursive linkage)
 
* Main process :
    * generate distance matrix,
    * initialize clusters with single point,
    * find 2 closests clusters, merge them in a new cluster, saving in the cluster the distance between the clusters that were merged
    * until there is a single cluster
 
* Exploit output
  * simplify cluster structure for memoruy efficiency ( single slice contains all object ids, cluster is a sub-slice of the table)
  * produce dentogram, with intercluster distance display ? 
  * allocate a new, unknown file to the closest cluster ?
