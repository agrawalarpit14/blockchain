from django.shortcuts import render, redirect
from django.contrib.auth.decorators import login_required
from django.contrib import messages
import subprocess

# Create your views here.
def welcome(request):
	return render(request, 'kyc/index.html')

@login_required
def webcam(request):
	return render(request, 'kyc/webcam.html')

@login_required
def terminal(request, key=0):
	states = [('KYC', '', 1),
			 ('Get aadhar number', '', 2),
			 ('Update KYC', '', 3),
			]
	me = subprocess.check_output("ls")
	if key != 0:
		if key==3:
			return redirect('webcam')
		if key==1:
			me = subprocess.check_output(['node', 'invoke.js'])
		if key==2:
			me = subprocess.check_output(['node', 'query.js'])
		messages.info(request, me.decode("utf-8"))
		
	context ={
		'names' : states
	}
	return render(request, 'kyc/terminal.html', context)