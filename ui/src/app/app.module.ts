import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AngularSplitModule } from 'angular-split';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'; 

import { AppComponent } from './app.component';
import { MaterialModule } from './material.module';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MaterialModule,
    AngularSplitModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
