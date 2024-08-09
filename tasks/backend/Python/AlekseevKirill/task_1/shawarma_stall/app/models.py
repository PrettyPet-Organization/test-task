from django.db import models


class Product(models.Model):
    """Модель Продукта"""
    title = models.CharField(max_length=256)
    price = models.DecimalField(max_digits=10, decimal_places=2)

    def __str__(self):
        return f"{self.title} - {self.price}"
    


class Order(models.Model):
    """Модель заказа"""

    status_choices = (
        ('in_processing', 'в обработке'),
        ('confirmed', 'подтвержден'),
        ('canceled', 'отменен'),
        ('made', 'сделан'),
    )

    product = models.ForeignKey(Product, on_delete=models.CASCADE,)
    status = models.CharField(max_length=20, choices=status_choices, default='in_processing')
    telephone = models.CharField(max_length=11)
    count = models.PositiveIntegerField(default=1)

    def __str__(self):
        return f"{self.telephone} - {self.status}"