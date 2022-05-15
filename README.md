# Go Generics

## Pengenalan Generic
- Generic adalah sebuah kemampuan menambahkan tipe parameter saat kita membuat function
- Berbeda dengan tipe data yang biasanya kita gunakan di function, generic memungkinkan kita bisa ubah bentuk tipe data sesuai dengan yang kita inginkan
- Fitu generic ada sejak Go versi 1.18

## Manfaat Generic
- Pengecekan ketika proses kompilasi
- Tidak perlu manual menggunakan pengecekan tipe dan konversi tipe data
- Memudahkan programmer membuat kode program yang generic sehingga bisa digunakan oleh berbagai tipe data

## Type Parameter
- Untuk menandai sebuah function ayng merupaka generic, kita perlu menambahkan Tipe Parameter pada function tersebut
- Pembuatan Type Parameter menggunakan tanda kurung kotak **[]**, di mana di dalam kurung kotak tersebut, kita tentukan nama Type Parameternya
- Hampir sama dengan di bahasa pemrograman lain. Best practicenya Type Parameter hanya menggunakan satu huruf saja seperti T, K, V, dll.

## Type Constraint
- Di bahasa pemrograman seperti Java, C#, dll, Type Parameter biasanya tidak perlu kita tentukan tipe datanya. Berbeda dengan di Go
- Dari pengalaman yang dilakukan para developer Go, akhirnya di Go, Type Parameter wajib memiliki constraint
- Type Constraint merupakan aturan yang digunakan untuk menentukan tipe data yang diperbolehkan pada Type Parameter
- Contoh, jika kita ingin Type Parameter bisa digunakan untuk semua tipe data, kita bisa gunakan interface{} kosong sebagai constraint-nya


