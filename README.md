# Magicsoft Compare Folder
Penjelasan Implementasi Compare folder test 

1. Membuat fungsi **IOReadir** untuk mendapatkan list dari direktory source dan target 
2. Membuat fungsi **compareFolder** untuk compare list file dari hasil fungsi **IOReadir**
    * membuat map untuk menampung nama file dan diberi nilai integer 
    * mengambil nama string file yang sama dari source dan target 
    * setiap nama string yang sama akan di tambahkan nilai integernya
    * buat statement ketika nilai int == 1 print file yang tidak ada di source tapi ada di target diberikan keterangan "DELETED"
    
