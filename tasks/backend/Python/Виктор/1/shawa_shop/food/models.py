from django.db import models


# Create your models here.

class Product(models.Model):
    CATEGORY_CHOICES = [
        ('none', 'Без категории'),
        ('shawarma', 'Шаурма'),
        ('beer', 'Пивко'),
    ]
    name = models.CharField(max_length=150, )
    price = models.IntegerField()
    description = models.TextField(blank=True)  # Здесь можно написать ингредиенты
    category = models.CharField(max_length=10, choices=CATEGORY_CHOICES, default='none')  # По дефолту - без категории

    def __str__(self):
        return f'{self.name} - {self.price}'
