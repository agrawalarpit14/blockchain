from django import forms
from django.contrib.auth.models import User
# from django.contrib.auth.forms import UserCreationForm
from .models import Profile, AdFile
from django.forms.widgets import TextInput
from django.forms import ModelForm
from django.contrib.auth.forms import UserCreationForm


# from .models import UserProfile
class UserRegisterForm(UserCreationForm):
    first_name = forms.CharField(widget=forms.TextInput())
    last_name = forms.CharField(widget=forms.TextInput())
    email = forms.EmailField()


    class Meta:
        model = User
        fields = ['first_name', 'last_name', 'username', 'email', 'password1', 'password2']

class UserUpdateForm(forms.ModelForm):
    email = forms.EmailField()

    class Meta:
        model = User
        fields = ['username', 'email']


class ProfileUpdateForm(forms.ModelForm):
    class Meta:
        model = Profile
        fields = ['image']


class AdForm(ModelForm):
    class Meta:
        model = AdFile
        fields = ['document', 'link', 'text', 'text_color', 'bg_color']
        widgets = {
            'text_color': TextInput(attrs={'type': 'color'}),
            'bg_color': TextInput(attrs={'type': 'color'}),
        }

    

