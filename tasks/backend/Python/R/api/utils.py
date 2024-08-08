from django.db import transaction


def create_order(serializer):
    try:
        vd = serializer.validated_data
        product = vd['product']
        order_quantity = vd['quantity']

        with transaction.atomic():
            product.quantity -= order_quantity
            product.save()
            order = serializer.save()

        return order

    except Exception as e:
        raise ValueError(f"Ошибка при создании заказа: {e}")
