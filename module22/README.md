# Домашнее задание по работе с терминалом linux

## Задание 1
Создайте директории, группы пользователей и нужных пользователей. Исходя из условий задания, решите, 
в какой группе должен быть создаваемый пользователь (и должен ли вообще), а также выдайте права так, 
чтобы выполнялись следующие пункты:
- Есть пользователь «Администратор» (admin). Он имеет полный доступ ко всем папкам.
- Есть группа пользователей «Бухгалтерия» (bookkeepers), пользователи в ней имеют полный доступ к директории 
invoices и доступы на чтение к директории web и директории отдела снабжения (работа со складом) store.
- Есть группа пользователей-кладовщиков (storekeepers), пользователи этой группы имеют полный доступ к папке 
store и доступ на чтение к папкам web и invoices.
- Есть группа пользователей-программистов (developers), пользователи этой группы имеют полный доступ к папке 
web и доступ на чтение к папкам invoices и store.

Все директории должны находится в /var/data.

```bash
# Создание директорий
sudo mkdir -p /var/data/{invoices,store,web}
# Создание групп
sudo groupadd bookkeepers && sudo groupadd storekeepers && sudo groupadd developers
# Создание админа
sudo useradd -m -s /bin/bash admin && sudo usermod -aG bookkeepers,storekeepers,developers admin
# Создание пользователей
sudo useradd -m -s /bin/bash -G bookkeepers user_bookkeeper
sudo useradd -m -s /bin/bash -G storekeepers user_storekeeper
sudo useradd -m -s /bin/bash -G developers user_developer
# Назначение владельцев и групп для директорий
sudo chown -R admin:bookkeepers /var/data/invoices
sudo chown -R admin:storekeepers /var/data/store
sudo chown -R admin:developers /var/data/web
# Настройка прав админу
sudo chmod -R 770 /var/data/invoices
sudo chmod -R 770 /var/data/store
sudo chmod -R 770 /var/data/web
# Настройка для invoices
sudo chmod -R 770 /var/data/invoices
sudo setfacl -m g:storekeepers:rx /var/data/invoices
sudo setfacl -m g:developers:rx /var/data/invoices
# Настройка для store
sudo chmod -R 770 /var/data/store
sudo setfacl -m g:bookkeepers:rx /var/data/store
sudo setfacl -m g:developers:rx /var/data/store
# Настройка для web
sudo chmod -R 770 /var/data/web
sudo setfacl -m g:bookkeepers:rx /var/data/web
sudo setfacl -m g:storekeepers:rx /var/data/web
# Наследование групп для новых файлов
sudo chmod g+s /var/data/invoices
sudo chmod g+s /var/data/store
sudo chmod g+s /var/data/web
```