name="object_storage"
go build -o ${name}.out
mkdir -p output
mv ${name}.out output/
cp conf/*.* output/