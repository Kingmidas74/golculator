import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AngularSplitModule } from 'angular-split';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'; 

import { AppComponent } from './app.component';
import { MaterialModule } from './material.module';
import { HttpClientModule } from '@angular/common/http';
import { HighlightModule, HIGHLIGHT_OPTIONS } from 'ngx-highlightjs';
import { AceDirective } from './ace.directive';
import { ModalCreateComponent } from './modal-create/modal-create.component';

@NgModule({
  declarations: [
    AppComponent,
    AceDirective,
    ModalCreateComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MaterialModule,
    AngularSplitModule,
    HttpClientModule,
    HighlightModule
  ],
  providers: [
    {
      provide: HIGHLIGHT_OPTIONS,
      useValue: {
        coreLibraryLoader: () => import('highlight.js/lib/core'),
        lineNumbersLoader: () => import('highlightjs-line-numbers.js'),
        languages: {
          lua: () => import('highlight.js/lib/languages/lua')
        }
      }
    }
  ],
  entryComponents:[ModalCreateComponent],
  bootstrap: [AppComponent]
})
export class AppModule { }
