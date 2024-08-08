from django.db import models


class ProductModel(models.Model):
    """
    Модель продукта.
    """
    name = models.CharField(max_length=200, verbose_name='Наименование продукта')
    price = models.DecimalField(max_digits=6, decimal_places=2, verbose_name='Цена в рублях')
    image = models.ImageField(verbose_name='Изображение')

    def __str__(self):
        return f'{self.name} - {self.price}'

    class Meta:
        verbose_name = 'Продукт'
        verbose_name_plural = 'Продукты'