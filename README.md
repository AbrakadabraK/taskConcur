1. Задача: Пинг-Понг (done :+1: )
   
Напишите программу, которая симулирует игру в пинг-понг между двумя горутинами. Используйте каналы для передачи «мяча» между двумя горутинами. Программа должна завершиться после 10 передач мяча.

Подсказка:

 - Используйте два канала: один для передачи "пинга" и другой для передачи "понга".

2. Задача: Пример с таймером и каналами (done :+1: )
   
Напишите программу, которая создает несколько горутин, каждая из которых будет выполнять работу с разным временем задержки. Каждая горутина должна отправлять результаты своей работы в канал, а основная программа должна выводить результаты по мере их поступления.

Подсказка:

-Используйте time.Sleep для моделирования задержек и select для чтения из канала.

3. Задача: Сумма элементов массива параллельно (done :+1: )
   
Разбейте массив чисел на несколько частей и напишите программу, которая вычисляет сумму элементов массива параллельно с использованием нескольких горутин.

Подсказка:

- Разделите массив на части, запустите горутины для суммирования каждой части и используйте канал для возврата результатов.
  
4. Задача: Пул воркеров (done :+1: )
   
Реализуйте пул воркеров. Программа должна запускать несколько горутин, каждая из которых будет обрабатывать задания, переданные через канал. Основная программа должна передавать задачи в канал и получать результаты.

Подсказка:

- Используйте буферизированный канал для передачи задач и другой канал для получения результатов.
  
5. Задача: Философы и вилка (done :+1: )
   
Классическая задача об обедающих философах. Реализуйте решение с использованием горутин и каналов, чтобы избежать взаимоблокировки (deadlock) при доступе к ресурсам (вилки).

Описание задачи:
Представьте, что за круглым столом сидят N философов. Каждый философ чередует два состояния: думает или ест. Для того чтобы поесть, философу необходимо одновременно взять две вилки — одну в левую руку и одну в правую (вилки расположены между философами). После еды философ кладёт вилки обратно на стол и снова начинает думать.

Главное ограничение: философ может использовать только две вилки, находящиеся рядом с ним.

Проблемы, которые нужно решить:
Мертвые блокировки (deadlock): Если все философы одновременно возьмут одну вилку (например, правую), они будут ждать, пока освободится вторая вилка, но ни один философ не сможет её взять, так как все уже захватили одну вилку.

Голодание (starvation): Даже если мёртвых блокировок удалось избежать, могут быть ситуации, когда один из философов постоянно оказывается последним, кто получает вилки, и, следовательно, никогда не ест.

Потокобезопасность и гонки: Необходимо организовать доступ к вилкам таким образом, чтобы философы не пытались одновременно взять одну и ту же вилку, создавая конфликт.

Подсказка:

- Можете использовать мьютексы или каналы для синхронизации доступа к вилкам.

6. Задача: Concurrent Map (done :+1: )
   
Реализуйте потокобезопасную карту (map) с использованием примитивов синхронизации. Программа должна позволять конкурентное чтение и запись в карту без гонок данных.

Подсказка:

- Используйте мьютексы (sync.Mutex) или RWMutex для управления доступом к данным.
  
7. Задача: Producer-Consumer (done :+1: )
   
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


11. Есть набор урлов. (done :+1: )


package main

func main() {
  var urls = []string{
    "http://ozon.ru",
    "https://ozon.ru",
    "http://google.com",
    "http://somesite.com",
    "http://non-existent.domain.tld",
    "https://ya.ru",
    "http://ya.ru",
    "http://ёёёё",
  }
}
Напишите программу, которая:

1. Поочередно выполнит http запросы по предложенному списку ссылок
в случае получения http-кода ответа на запрос "200 OK" печатаем на экране "адрес url - ok"
в случае получения http-кода ответа на запрос отличного от "200 OK" либо в случае ошибки печатаем на экране "адрес url - not ok"

2. Модифицируйте программу таким образом, чтобы использовались каналы для коммуникации основного потока с горутинами. Пример:
Запросы по списку выполняются в горутинах.
Печать результатов на экран происходит в основном потоке

3. Модифицируйте программу таким образом, чтобы нигде не использовалась длина слайса урлов. Считайте, что урлы приходят из внешнего источника. Сколько их будет заранее - неизвестно. Предложите идиоматичный вариант, как ваша программа будет узнавать об окончании списка и передавать сигнал об окончании действий далее.
4. (необязательно, можно обсудить устно, чтобы убедиться, что кандидат понимает идею контекста, либо предложить как домашнее задание) Модифицируйте программу таким образом, что бы при получении 2 первых ответов с "200 OK" остальные запросы штатно прерывались.
При этом необходимо напечатать на экране сообщение о завершении запроса.
5. (необязательно, можно обсудить устно) Предложите отрефакторить код. Какие тесты кандидат написал бы к этому коду?
Предложите написать код теста и интерфейсы, для которых будут генериться моки. (Как показывает практика это самая сложная часть задачи)



12. (done :+1: )
Нужно написать простую библиотеку in-memory cache.
Для простоты считаем, что у нас бесконечная память и нам не нужно задумываться об удалении ключей из него.
Реализация должна удовлетворять интерфейсу:


type Cache interface {
    Set(k, v string)
    Get(k string) (v string, ok bool)
}
