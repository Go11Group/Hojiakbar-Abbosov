-- UPDATE operatori yordamida ma'lumotlarni o'zgartirishimiz mumkin.
-- Bu operator bilan bir nechta qatorlarni o'zgartirishimiz mumkin.
-- Misol uchun, Buyurtmalar jadvalidagi narxlarni 20% ga oshirish:

UPDATE Buyurtmalar SET Narxi = Narxi * 1.2;

-- DELETE operatori yordamida ma'lumotlarni databazadan o'chirishimiz mumkin. 
-- Misol uchun, quyidagi Buyurtmalar jadvalidagi "Olma" buyurtmasini o'chiramiz:

DELETE FROM Buyurtmalar WHERE Ismi = 'Olma';

--GROUP BY operatori yordamida ma'lumotlarni guruhlab, 
--				aggregat funksiyalarni (masalan, SUM, AVG) ishlatishimiz mumkin.
-- Misol uchun, Buyurtmalar jadvalidagi miqdorlar bo'yicha guruhlab, umumiy summasini topamiz:

SELECT Ismi, SUM(Miqdori) AS UmumiyMiqdor FROM Buyurtmalar GROUP BY Ismi;

-- ORDER BY operatori yordamida ma'lumotlarni saralashimiz mumkin.
-- Misol uchun, Buyurtmalar jadvalidagi buyurtmalarni narx bo'yicha kamdan ko'p'gacha saralash:


SELECT * FROM Buyurtmalar ORDER BY Narxi ASC;

-- Join (4 ta joinlar):
-- Join operatorlari yordamida bir-biriga bog'liq jadvallar orasida bog'lanishlar qilishimiz mumkin.

-- INNER JOIN: ikkita jadvalni bog'lash. 
-- Masalan, Buyurtmalar va Mahsulotlar jadvallari o'rtasida bog'lanish:

SELECT Buyurtmalar.Ismi, Mahsulotlar.Nomi
FROM Buyurtmalar
INNER JOIN Mahsulotlar ON Buyurtmalar.MahsulotID = Mahsulotlar.ID;

-- LEFT JOIN: Birinchi jadvalni o'ngdan bog'lash. 
-- Masalan, Buyurtmalar jadvali bilan Mahsulotlar jadvallarini bog'lash:

SELECT Buyurtmalar.Ismi, Mahsulotlar.Nomi
FROM Buyurtmalar
LEFT JOIN Mahsulotlar ON Buyurtmalar.MahsulotID = Mahsulotlar.ID;

-- RIGHT JOIN: Ikkinchi jadvalni o'ngdan bog'lash. 
-- Masalan, Mahsulotlar jadvali bilan Buyurtmalar jadvallarini bog'lash:


SELECT Buyurtmalar.Ismi, Mahsulotlar.Nomi
FROM Buyurtmalar
RIGHT JOIN Mahsulotlar ON Buyurtmalar.MahsulotID = Mahsulotlar.ID;

-- FULL JOIN: Ikkita jadvallarni to'la bog'lash. 
-- Masalan, Buyurtmalar va Mahsulotlar jadvallarini to'la bog'lash:

SELECT Buyurtmalar.Ismi, Mahsulotlar.Nomi
FROM Buyurtmalar
FULL JOIN Mahsulotlar ON Buyurtmalar.MahsulotID = Mahsulotlar.ID;

