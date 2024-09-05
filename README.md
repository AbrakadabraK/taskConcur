1. Задача: Пинг-Понг
Напишите программу, которая симулирует игру в пинг-понг между двумя горутинами. Используйте каналы для передачи «мяча» между двумя горутинами. Программа должна завершиться после 10 передач мяча.

Подсказка:

 - Используйте два канала: один для передачи "пинга" и другой для передачи "понга".

2. Задача: Пример с таймером и каналами
Напишите программу, которая создает несколько горутин, каждая из которых будет выполнять работу с разным временем задержки. Каждая горутина должна отправлять результаты своей работы в канал, а основная программа должна выводить результаты по мере их поступления.

Подсказка:

-Используйте time.Sleep для моделирования задержек и select для чтения из канала.

3. Задача: Сумма элементов массива параллельно
Разбейте массив чисел на несколько частей и напишите программу, которая вычисляет сумму элементов массива параллельно с использованием нескольких горутин.

Подсказка:

- Разделите массив на части, запустите горутины для суммирования каждой части и используйте канал для возврата результатов.
  
4. Задача: Пул воркеров
Реализуйте пул воркеров. Программа должна запускать несколько горутин, каждая из которых будет обрабатывать задания, переданные через канал. Основная программа должна передавать задачи в канал и получать результаты.

Подсказка:

- Используйте буферизированный канал для передачи задач и другой канал для получения результатов.
  
5. Задача: Философы и вилка
Классическая задача об обедающих философах. Реализуйте решение с использованием горутин и каналов, чтобы избежать взаимоблокировки (deadlock) при доступе к ресурсам (вилки).

Подсказка:

Можете использовать мьютексы или каналы для синхронизации доступа к вилкам.
6. Задача: Concurrent Map
Реализуйте потокобезопасную карту (map) с использованием примитивов синхронизации. Программа должна позволять конкурентное чтение и запись в карту без гонок данных.

Подсказка:

- Используйте мьютексы (sync.Mutex) или RWMutex для управления доступом к данным.
  
7. Задача: Producer-Consumer
Напишите программу, которая симулирует взаимодействие производителя и потребителя. Производитель генерирует данные и отправляет их в канал, а потребитель получает данные из канала и обрабатывает их.

Подсказка:

- Используйте буферизированный канал для управления потоком данных.
  
8. Задача: Параллельная сортировка (Merge Sort)
Реализуйте параллельный алгоритм сортировки слиянием (Merge Sort) с использованием горутин. Разделите массив на подмассивы и сортируйте каждый подмассив в отдельной горутине.

Подсказка:

 - Рекурсивно разбивайте массив на части и используйте горутины для сортировки каждой части.
   
9. Задача: Программа с таймаутом
Напишите программу, которая запускает длительную операцию в горутине. Если операция не завершится за определенное время, программа должна завершить выполнение с сообщением о таймауте.

Подсказка:

- Используйте time.After и select, чтобы отслеживать время выполнения операции.
  
10. Задача: Barbershop (Проблема парикмахера)
Реализуйте решение проблемы парикмахера, где есть ограниченное количество стульев, и клиенты должны либо подождать, либо уйти, если все стулья заняты. Парикмахер работает только тогда, когда есть клиенты.

Подсказка:

- Используйте горутины для симуляции клиентов и парикмахера, а каналы для управления очередью клиентов.

Эти задачи помогут освоить параллелизм и синхронизацию в Go.
