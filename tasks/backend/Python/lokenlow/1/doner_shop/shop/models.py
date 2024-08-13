from django.db import models
from django.utils.text import slugify
from django.core.validators import MinValueValidator

from decimal import Decimal


class Good(models.Model):

    name = models.CharField(max_length=100, null=False)
    price = models.DecimalField(max_digits=7, decimal_places=2, default=Decimal('0.00'), null=False,
                                validators=[MinValueValidator(Decimal('0.00'))])
    stock = models.PositiveIntegerField(null=False, default=0)
    description = models.TextField(max_length=2000, null=True, blank=True, default=None)

    def __str__(self):
        return self.name

    @property
    def slug(self):
        return slugify(self.name)

    def save(self, *args, **kwargs):
        super().save(*args, **kwargs)
