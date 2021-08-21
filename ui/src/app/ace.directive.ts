import { Directive, ElementRef, EventEmitter } from "@angular/core";

import * as ace from "ace-builds";

@Directive({
    selector: 'ace-editor',
    inputs: [
        "text"
    ],
    outputs: [
        "textChanged"
    ]
})
export class AceDirective { 
    private editor;
    public textChanged: EventEmitter<string>;

    /**
     * Sets the editor's text.
     */
    set text(s: string) {
        this.editor.setValue(s);
        this.editor.clearSelection();
        this.editor.focus();
    }
    
    constructor(elementRef: ElementRef) {
        this.textChanged = new EventEmitter<string>();
        ace.config.set("fontSize", "21px");
        ace.config.set('basePath', 'https://unpkg.com/ace-builds@1.4.12/src-noconflict');
        // this is the <div ace-editor> root element
        let el = elementRef.nativeElement;
        el.classList.add("editor");
        el.style.height = "29vh";
        el.style.width = "100%";
        this.editor = ace.edit(el);
        //this.editor.resize(true);
        /*  ace.config.set("fontSize", "14px");
    ace.config.set('basePath', 'https://unpkg.com/ace-builds@1.4.12/src-noconflict');

    const aceEditor = ace.edit(this.editor.nativeElement);
    aceEditor.session.setValue("");
    aceEditor.setTheme('ace/theme/twilight');
    aceEditor.session.setMode('ace/mode/lua');

    aceEditor.on("change", () => {
      console.log(aceEditor.getValue());
    });*/

        this.editor.setTheme("ace/theme/twilight");
        this.editor.getSession().setMode("ace/mode/lua");
        
        this.editor.on("change", (e) => {
            // discard the delta (e), and provide whole document
            this.textChanged.next(this.editor.getValue());
        });
    }
}