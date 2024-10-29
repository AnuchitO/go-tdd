# go-tdd
Learing Test-Driven Development with Go programming language



## Requirements

## Tuk Tuk Invoice Requirements

### 1: Fare Calculation:
- The system must determine the total fare for a ride based on the distance and waiting time, utilizing the
- formula: ฿4 * kilometer + ฿1 * minute of waiting
- The distance should be rounded up to the nearest 0.5 km, and waiting time should be rounded up to the nearest minute.
### 2: Minimum Fare Application:
- The system is required to impose a minimum fare of 35 baht if the calculated fare falls below this threshold.
### 3: Total Fare Calculation for Multiple Rides:
- For a given list of multiple rides, the system must compute the overall total fare.
- The fare calculation and minimum fare rules should be applied individually to each ride in the list.
### 4: Invoice Details Printing:
- The system should generate a detailed invoice that includes:

```
Invoice:
    Ride Fare: ฿51.00 for Distance: 12.0 km, Waiting Time: 3 minutes
    Ride Fare: ฿35.00 for Distance: 7.5 km, Waiting Time: 3 minutes
    Ride Fare: ฿37.00 for Distance: 8.5 km, Waiting Time: 3 minutes
Total Number of Rides: 3
Total Fare: ฿123.00
Average Fare Per Ride: ฿41.00
```

## ความต้องการของระบบคิดค่าโดยสารรถตุ๊กตุ๊ก

### 1: การคำนวณค่าโดยสาร:
- ระบบจะต้องคำนวณค่าโดยสารรวมทั้งหมดของการเดินทาง
- โดยใช้สูตร: ฿4 * กิโลเมตร + ฿1 * นาทีที่รอ
- ระบบจะต้องปัดเศษของระยะทางที่เหลืออยู่ให้เป็น 0.5 กิโลเมตร และปัดเศษของเวลาที่รออยู่ให้เป็นนาที
### 2: การใช้ค่าโดยสารขั้นต่ำ:
- ระบบจะต้องคิดค่าโดยสารขั้นต่ำ 35 บาท หากค่าโดยสารที่คำนวณได้น้อยกว่านี้
### 3: การคำนวณค่าโดยสารรวมสำหรับการเดินทางหลายรอบ:
- สำหรับรายการการเดินทางหลายรอบ ระบบจะต้องคำนวณค่าโดยสารรวมทั้งหมด
- ระบบจะต้องคำนวณค่าโดยสารและใช้กฎขั้นต่ำแยกตามการเดินทางแต่ละรอบ
### 4: การพิมพ์รายละเอียดใบแจ้งหนี้:
- ระบบจะต้องสร้างใบแจ้งหนี้ที่มีรายละเอียดอย่างละเอียด ซึ่งรวมถึง:
```
Invoice:
    Ride Fare: ฿51.00 for Distance: 12.0 km, Waiting Time: 3 minutes
    Ride Fare: ฿35.00 for Distance: 7.5 km, Waiting Time: 3 minutes
    Ride Fare: ฿37.00 for Distance: 8.5 km, Waiting Time: 3 minutes
Total Number of Rides: 3
Total Fare: ฿123.00
Average Fare Per Ride: ฿41.00
```
