# Generated by Django 2.1.2 on 2018-12-27 12:52

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
        ('users', '0002_auto_20181220_1155'),
    ]

    operations = [
        migrations.CreateModel(
            name='AdFile',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('document', models.FileField(blank=True, upload_to='customer_files')),
                ('link', models.CharField(blank=True, default='', max_length=100)),
                ('date_uploaded', models.DateTimeField(auto_now_add=True)),
                ('text', models.CharField(blank=True, default='', max_length=100)),
                ('text_color', models.CharField(blank=True, default='', max_length=100)),
                ('bg_color', models.CharField(blank=True, default='', max_length=100)),
                ('task_by', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL)),
            ],
        ),
    ]
