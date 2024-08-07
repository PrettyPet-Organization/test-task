from django.urls import path
from .views import FormuleView


app_name = 'Formule'

urlpatterns = [
    path('', FormuleView, name='formule_view'),

]
