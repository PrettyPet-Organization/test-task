from rest_framework import serializers
from .models import Sklad


class SkladSerializer(serializers.ModelSerializer):
    class Meta:
        model = Sklad
        fields = '__all__'










