#!/bin/bash
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, New World!} http://localhost:8888/files/new.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, New World!"}' http://localhost:8888/files/new.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, Old World!} http://localhost:8888/files/old.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, Old World!"}' http://localhost:8888/files/old.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, Baby World!} http://localhost:8888/files/baby.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, Baby World!"}' http://localhost:8888/files/baby.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, Girl World!} http://localhost:8888/files/girl.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, Girl World!"}' http://localhost:8888/files/girl.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, Boy World!} http://localhost:8888/files/boy.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, Boy World!"}' http://localhost:8888/files/boy.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, Female World!} http://localhost:8888/files/female.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, Female World!"}' http://localhost:8888/files/female.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, Male World!} http://localhost:8888/files/male.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, Male World!"}' http://localhost:8888/files/male.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, Male World!} http://localhost:8888/files/cat.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, Male World!"}' http://localhost:8888/files/cat.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X POST -H Content-Type: application/json -d {data:Hello, Male World!} http://localhost:8888/files/dog.txt"
curl -X POST -H "Content-Type: application/json" -d '{"data":"Hello, Male World!"}' http://localhost:8888/files/dog.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl  http://localhost:8888/files/new.txt"
curl  http://localhost:8888/files/new.txt
echo "---"
echo "curl  http://localhost:8888/files/old.txt"
curl  http://localhost:8888/files/old.txt
echo "---"
echo "curl  http://localhost:8888/files/baby.txt"
curl  http://localhost:8888/files/baby.txt
echo "---"
echo "curl  http://localhost:8888/files/cat.txt"
curl  http://localhost:8888/files/cat.txt
echo "---"
echo "curl  http://localhost:8888/files/boy.txt"
curl  http://localhost:8888/files/boy.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X DELETE  http://localhost:8888/files/new.txt"
curl -X DELETE  http://localhost:8888/files/new.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X DELETE  http://localhost:8888/files/old.txt"
curl -X DELETE  http://localhost:8888/files/old.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X DELETE  http://localhost:8888/files/male.txt"
curl -X DELETE  http://localhost:8888/files/male.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
echo "curl -X DELETE  http://localhost:8888/files/boy.txt"
curl -X DELETE  http://localhost:8888/files/boy.txt
echo "---"
echo "curl  http://localhost:8888/files"
curl  http://localhost:8888/files
echo "---"
