# Context based search engine
Fuzzy searching based on annotated and related documents
___

<br />

- [X] Fuzzy search
- [X] Elasticsearch
- [X] Annotation/classification
- [ ] BigQuery
- [ ] Hadoop

<br />
<br />

### Instructions to run

<br />

#### Setup elastic


```bash
$ chmod +x test_elastic
$ docker-compose up
```

<br />

#### Test elastic connection

```bash
$ ./test_elastic
```

#### Run main server

```bash
$ go run main.go
```

<br />


### Benchmarks on localhost

<br />

| Operation performed |  Data involved | Time taken |
|:---:|:---:|:---:|
| Extract all text | 20 files (PDFs), 66799 lines of text | 1538 ms |
| Search for given key | Query sent to stored data on elasticsearch | 6 ms |

<br />
<br />


| Service name | Role | PORT |
|:---:|:---:|:---:|
| gorest | Main text extraction/annotation/search API | 3000 |
| elasticsearch | Mount and fuzzy search data | 9200 |

<br />
<br />


![Workflow](static/img/workflow.png)