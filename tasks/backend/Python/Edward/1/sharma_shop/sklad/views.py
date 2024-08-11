from django.shortcuts import render
from rest_framework import generics
from rest_framework.views import APIView
from rest_framework.response import Response
from django.forms import model_to_dict

from .models import Sklad
from .serializers import SkladSerializer


class SkladAPIView(APIView):
    def get(self, request):
        lst = Sklad.objects.all()
        return Response({'post' : SkladSerializer(lst, many=True).data})

    def post(self,request):
        serializer = SkladSerializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        serializer.save()

        return Response({'new_produkt': serializer.data})

    def put(self, request, *args, **kwargs):
        pk = kwargs.get('pk', None)
        if not pk:
            return Response({'error': 'Method PUT not allowed'})

        try:
            instance = Sklad.objects.get(pk=pk)
        except:
            return Response({'error' : 'Object does not exists'})

        serializer = SkladSerializer(data=request.data, instance=instance)
        serializer.is_valid(raise_exception=True)
        serializer.save()
        return Response({'post': serializer.data})

