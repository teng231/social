### Quy tắc 1
* Chỉ có core được log. Các phần khác không được log trừ trường hợp đặc biệt là chạy config
* Các thành phần khác ko đươc hiểu nhau

* Core không cân hiểu các phần khác nhau của hệ thống mà hiểu các ánh xạ.

* Core luôn có các kết nối tới các phần bằng drive

* Logic ko được để ngoài core trừ trường hợp đặc hữu dành riêng cho drive Vd: bson.NewObjectID() là thành phần của mongo

