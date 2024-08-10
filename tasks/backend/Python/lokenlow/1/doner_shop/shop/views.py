from django.shortcuts import get_object_or_404
from rest_framework import permissions, viewsets, status
from rest_framework.decorators import action
from rest_framework.response import Response

from .models import Good
from .serializers import GoodsSerializer


class GoodsViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows goods to be viewed or edited.
    """
    queryset = Good.objects.all().order_by('name')
    serializer_class = GoodsSerializer
    permission_classes = [permissions.IsAuthenticated]

    @action(detail=True, methods=['delete'], url_path='delete/(?P<id>[^/.]+)')
    def delete(self, request, pk=None):
        if not pk:
            return Response('Bulk deletions are not supported', status=status.HTTP_400_BAD_REQUEST)
        good = get_object_or_404(Good, id=pk)
        good.delete()
        return Response(status=status.HTTP_204_NO_CONTENT)
