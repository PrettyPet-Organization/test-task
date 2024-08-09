from django.db import models


class Ingredient(models.Model):
    name = models.CharField(verbose_name='Название')
    measurement_unit = models.CharField(verbose_name='Единицы измерения')
    
    class Meta:
        verbose_name = 'Ингредиент'
        verbose_name_plural = 'Ингредиенты'

    def __str__(self):
        return f'Ингредиент {self.name}.'


class Beer(models.Model):
    name = models.CharField(verbose_name='Название')
    sort = models.CharField(verbose_name='Сорт')
    category = models.CharField(verbose_name='Категория')
    alcohol = models.IntegerField(verbose_name='Крепость')
    density = models.IntegerField(verbose_name='Плотность')
    producer = models.CharField(verbose_name='Производитель')
    price = models.IntegerField(verbose_name='Цена')
    
    class Meta:
        verbose_name = 'Пиво'
        verbose_name_plural = 'Пиво'
    
    def __str__(self):
        return f'Пиво {self.name} от {self.producer} за {self.price} рублей.'


class Shaurma(models.Model):
    name = models.CharField(verbose_name='Название')
    ingredients = models.ManyToManyField(Ingredient, through='Recipe', verbose_name='Ингредиенты', through_fields=('shaurma', 'ingredient'))
    
    class Meta:
        verbose_name = 'Шаурма'
        verbose_name_plural = 'Шаурма'
    
    def __str__(self):
        return f'Шаурма {self.name}.'


class Recipe(models.Model):
    shaurma = models.ForeignKey(Shaurma, verbose_name='Шаурма', on_delete=models.CASCADE)
    ingredient = models.ForeignKey(Ingredient, verbose_name='Ингредиент', on_delete=models.CASCADE)
    amount = models.IntegerField(verbose_name='Количество ингредиента')
    
    class Meta:
        verbose_name = 'Ингредиент в шаурме'
        verbose_name_plural = 'Ингредиенты в шаурме'
    
    def __str__(self):
        return f'{self.ingredient} в {self.shaurma}.'
