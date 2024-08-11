from django.db import models

# Create your models here.

class Sklad(models.Model):
    name_produkt = models.CharField(max_length=50)
    count = models.IntegerField(default=0)
    price = models.DecimalField(max_digits=6, decimal_places=2)

    def __str__(self):
        return self.name_produkt






