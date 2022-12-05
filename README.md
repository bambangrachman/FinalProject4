# Final Project 4 Kelompok 4 Aplikasi Toko Belanja
Ini adalah project keempat dari program MSIB di Hacktiv8. Project kali ini adalah membuat sebuah aplikasi berjudul "Toko Belanja" dimana terdapat seorang admin yang berwenang melakukan perintah CRUD pada category dan juga product dan customer-customer yang bisa melakukan top up untuk membeli product dan juga bisa melihat transaksi pembeliannya.

## Our Team
* Alessandro (Category)
* Alfin (Transaction History)
* Fitri (Product)
* Faiz (User)

### End Points
**PRODUCT**
NOTES : Selain dari perintah **'GET'**, semua perintah lain hanya bisa diakses oleh admin. Jika customer melakukan perintah 'POST','DEL','PUT', maka akses akan ditolak dengan response seperti ini :
    ```json
    {
        "error": "You aren't allowed to do this! You are not Admin!""
        }
        ```
* GET :
    * Untuk menampilkan semua product dapat dengan menggunakan url :
    `http://localhost:8080/products` dengan method **GET**
    * Output response yang dihasilkan adalah :
        ```json
        [
            {
		"id": 1,
		"CreatedAt": "2022-12-05T03:30:11.716Z",
		"UpdatedAt": "2022-12-05T03:30:11.716Z",
		"title": "tote bag",
		"stock": 5,
		"price": 25000,
		"category_id": 1
        },
        {
		"id": 2,
		"CreatedAt": "2022-12-05T03:30:57.85Z",
		"UpdatedAt": "2022-12-05T03:30:57.85Z",
		"title": "sling bag",
		"stock": 5,
		"price": 35000,
		"category_id": 1
        },
            ]
        ```

* POST :
    * Untuk menambahkan product baru dapat dengan menggunakan url :
    `http://localhost:8080/products` dengan method **POST**
    * Kemudian gunakan json berikut untuk membuat datanya:
        ```json
            {
                "title":"cute bag",
                "stock":5,
                "price":85000,
                "category_id":1
            }
        ```
    * Untuk akses endpointnya dibutuhkan request autorisasi token yang didapatkan dari response endpoint user/login. (**Hanya bisa diakses oleh user dengan role ADMIN**)
    * Output response yang dihasilkan adalah :
        ```json
        {
            "id": 7,
            "title": "cute bag",
            "stock": 5,
            "price": 85000,
            "category_id": 1,
            "created_at": "2022-12-05T14:05:39.712+07:00"
            }
        ```

* PUT :
    * Misalnya untuk mengedit data product dengan id 7 dapat dengan menggunakan url :
    `http://localhost:8080/products/7` dengan method **PUT**
    * Kemudian gunakan json berikut untuk mengedit datanya:
        ```json
        {
            "title":"cute tiny bag",
            "stock":5,
            "price":45000,
            "category_id":1
        }
        ```
    * Untuk akses endpointnya dibutuhkan request autorisasi token yang didapatkan dari response endpoint user/login. (**Hanya bisa diakses oleh user dengan role ADMIN**)
    * Output response yang dihasilkan adalah :
        ```json
        {
            "id": 7,
            "title": "cute tiny bag",
            "stock": 5,
            "price": 45000,
            "created_at": "2022-12-05T07:05:39.712Z",
            "updated_at": "2022-12-05T14:10:42.416+07:00"
        }
        ```
* DELETE :
    * Misalnya untuk menghapus product dengan id 7 dapat dengan menggunakan url :
    `http://localhost:8080/products/7` dengan method **DELETE**
    * Untuk akses endpointnya dibutuhkan request autorisasi token yang didapatkan dari response endpoint user/login. (**Hanya bisa diakses oleh user dengan role ADMIN**)
    * Output response yang dihasilkan adalah :
        ```json
        {
        "message": "Product has been successfully deleted"
        }
        ```

**TASK**
* GET :
    * Untuk menampilkan semua data task dapat dengan menggunakan url :
    `http://localhost:8080/task` atau `https://fp3-kelompok4.up.railway.app/task` dengan method **GET**
    * Untuk dapat mengakses endpointnya dibutuhkan autorisasi token yang didapatkan dari response endpoint user/login
    * Output response yang dihasilkan adalah :
        ```json
        {
            "task": [
                {
                    "id": 4,
                    "title": "Menyelasaikan soal latihan MK Kalkulus",
                    "status": false,
                    "description": "Menyelesaikan soal latihan 1-10 BAB II",
                    "user_id": 3,
                    "category_id": 1,
                    "created_at": "2022-11-25T13:57:02Z",
                    "UpdatedAt": "2022-11-25T13:57:02Z",
                    "user": {
                        "id": 3,
                        "email": "kelompok4@gm.com",
                        "full_name": "kelompok4"
                    }
                }
            ]
        }
        ```

* POST :
    * Untuk membuat data task baru dapat dengan menggunakan url :
    `http://localhost:8080/task` atau `https://fp3-kelompok4.up.railway.app/task` dengan method **POST**
    * Kemudian gunakan json berikut untuk membuat datanya:
        ```json
        {
            "title" : "Menyelasaikan soal latihan MK Kimia",
            "description" : "Menyelesaikan soal latihan 5-10 BAB VII",
            "category_id" : 1
        }
        ```
    * Untuk akses endpointnya dibutuhkan request autorisasi token yang didapatkan dari response endpoint user/login.
    * Output response yang dihasilkan adalah :
        ```json
        {
            "id": 5,
            "title": "Menyelasaikan soal latihan MK Kimia",
            "status": false,
            "description": "Menyelesaikan soal latihan 5-10 BAB VII",
            "user_id": 3,
            "category_id": 1,
            "created_at": "2022-11-25T21:38:37.892+07:00",
            "updated_at": "2022-11-25T21:38:37.892+07:00"
        }
        ```
