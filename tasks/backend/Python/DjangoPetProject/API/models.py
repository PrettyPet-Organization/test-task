from django.db import models
from django.urls import reverse
from django.utils.translation import gettext_lazy as _


class Product(models.Model):
    """Класс продукта"""
    name = models.CharField(max_length=200, verbose_name=_('name'))
    price = models.FloatField(verbose_name=_('price'))
    count = models.IntegerField(null=True, blank=True, verbose_name=_('count'))
