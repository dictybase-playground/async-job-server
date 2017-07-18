# Async Job Server


Requirements:
1. Blastp installed https://blast.ncbi.nlm.nih.gov/Blast.cgi?CMD=Web&PAGE_TYPE=BlastDocs&DOC_TYPE=Download
2. Gearman server installed http://gearman.org/getting-started/
3. Download database http://dictybase.org/db/cgi-bin/dictyBase/download/blast_databases.pl
4. Create database with
 ```makeblastdb -in dicty_primary_protein -dbtype prot
```
Setup:
1. Start gearmand server ```/usr/local/sbin/gearmand -L 127.0.0.1 -p 4730 --verbose INFO```
2. Start worker

Parameters:
These are the paramaeters that should be sent via JSON to the worker:
```
type Arguments struct {
	Database string  `json:"database"`
	Query    string  `json:"query"`
	Evalue   float64 `json:"evalue"`
	Numalign int     `json:"numalign"`
	Wordsize int     `json:"wordsize"`
	Matrix   string  `json:"matrix"`
	Seg      bool    `json:"seg"`
	Gapped   bool    `json:"gapped"`
}```

An example:
```
a := &Arguments{
  Database: "dicty_primary_protein",
  Query:    "test_query.fsa",
  Evalue:   0.1,
  Numalign: 50,
  Wordsize: 3,
  Matrix:   "PAM30",
  Seg:      true,
  //Gapped:   false,
}
```
