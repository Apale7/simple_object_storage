name="object_storage"
go build -o ${name}.out
mkdir -p output
mv ${name}.out output/
cp -r static output/
cp conf/*.* output/