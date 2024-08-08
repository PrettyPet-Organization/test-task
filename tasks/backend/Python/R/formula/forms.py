from django import forms


class FormulaForm(forms.Form):
    choices = [
        ('add', 'Add'),
        ('rm', 'Remove'),
        ('mul', 'Multiply'),
        ('div', 'Divide'),
        ('mod', 'Modulus')
    ]

    operator = forms.ChoiceField(choices=choices)
    a = forms.IntegerField(required=True)
    b = forms.IntegerField(required=True)