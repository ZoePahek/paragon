// import { Monaco } from './monaco';
import * as monaco from 'monaco-editor';
import {BuiltIns} from './grammar';
import {Language, LanguageConfig} from './language';
import {Theme} from './theme';


export const LanguageID = 'renegade';

export const Register = () => {
  monaco.languages.register({id: LanguageID});

  monaco.editor.defineTheme(LanguageID, Theme);
  monaco.editor.setTheme(LanguageID);
  monaco.languages.setMonarchTokensProvider(LanguageID, Language);
  monaco.languages.setLanguageConfiguration(LanguageID, LanguageConfig(m));
  monaco.languages.registerCompletionItemProvider(LanguageID, {
    provideCompletionItems: (model, position) => {
      let word = model.getWordUntilPosition(position);
      let range = {
        startLineNumber: position.lineNumber,
        startColumn: word.startColumn,
        endLineNumber: position.lineNumber,
        endColumn: word.endColumn
      };
      return {
        suggestions: BuiltIns.map((func) => {
          return {
            label: func.name,
            detail: func.getDetail(),
            documentation: func.getDocs(),
            insertText: func.getInsertText(),
            insertTextRules:
                monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
            kind: monaco.languages.CompletionItemKind.Function,
            range: range,
            command: {
              id: 'editor.action.triggerParameterHint',
              title: 'editor.action.triggerParameterHint',
            }
          };
        }),
      }
    },
  });
  monaco.languages.registerSignatureHelpProvider(LanguageID, {
    signatureHelpTriggerCharacters: ['(', ','],
    // signatureHelpRetriggerCharacters: [','],
    provideSignatureHelp: (model, position, token, context) => {
      // Default value if no signatures are found
      let noSignatures = {
        value: {signatures: [], activeParameter: -1, activeSignature: -1},
        dispose: () => {},
      };
      let startPos = position;
      let endPos = position;

      // Start of function call
      let startMatch = model.findPreviousMatch(
          /\w+(?=\()/.source, position, true, true, null, true);
      if (!startMatch) {
        console.log('NO PREV FUNCTION', startPos, endPos, startMatch)
        return noSignatures;
      }
      startPos = startMatch.range.getEndPosition();

      // End of function call
      let endMatch =
          model.findNextMatch(/\)/.source, position, true, true, null, true);
      if (endMatch) {
        endPos = endMatch.range.getEndPosition();
      }

      // Out of function call bounds
      if (endPos.isBefore(position)) {
        console.log('POS OUT OF BOUNDS', startPos, endPos, startMatch, endMatch)
        return noSignatures;
      }

      // Find signature index
      let funcName = model.getWordUntilPosition(startPos).word;
      let sigIndex = BuiltIns.findIndex(({name}) => name === funcName);

      // No signature index found
      if (sigIndex < 0 || sigIndex >= BuiltIns.length) {
        console.log(
            'NO SIG INDEX', startPos, endPos, startMatch, endMatch, funcName,
            sigIndex);
        return noSignatures;
      }

      // Compute the signature
      let sig = BuiltIns[sigIndex].getSignature();

      // Get all param characters within the function call
      let funcBody = model.getValueInRange({
        startColumn: startPos.column,
        startLineNumber: startPos.lineNumber,
        endLineNumber: position.lineNumber,
        endColumn: position.column,
      })

      // Determine param index
      let funcBodyTokens =
          funcBody.match(/(?:(?<!['"])\b\w+|['"][^'"]*['"])\s*,/);
      let paramIndex = (!funcBodyTokens || funcBodyTokens.length <= 0) ?
          0 :
          funcBodyTokens.length >= sig.parameters.length ?
          sig.parameters.length - 1 :
          funcBodyTokens.length;

      console.log(
          'SIG FOUND', startPos, endPos, startMatch, endMatch, funcName,
          sigIndex, sig, funcBody, funcBodyTokens, paramIndex);
      return {
        value: {
          signatures: [sig],
          activeSignature: 0,
          activeParameter: paramIndex,
        },
        dispose: () => {},
      };
    },
  });
}
