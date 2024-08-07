from django import forms


class FormuleForm(forms.Form):
    """Форма для отображения формулы"""
    a = forms.IntegerField(required=True)
    b = forms.IntegerField(required=True)