Алгоритм следующий:
    Если длина haystack или needle равны 0, тогда возвращаем 0.
    Как находим совпадение по первому символу, запоминаем его, далее:
        * если совпали все символы, то возвращаем индекс первого элемента
        * если не совпали все символы, но совпала часть, то возвращаем обход на индекс первого совпавшего + 1
    Если алгоритм обошел всю строку, но не нашел совпадения, возвращаем -1.

