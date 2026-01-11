function getIndex(char: string): number {
    return char.charCodeAt(0) - 'a'.charCodeAt(0);
}


function countTotalSteps(tc:string): number {
    const splittedTestCase:string[] = tc.split(' ');
   const s1 = splittedTestCase[0]
   const s2 = splittedTestCase[1]

   let alphabetList:number[] = [];
   for(let i = 0; i < 26; i++) {
       alphabetList[i] = 0;
   }

   for(const char of s1) {
       alphabetList[getIndex(char)]++;
   }

   let totalSame = 0;
   for(const char of s2) {
       if(alphabetList[getIndex(char)] > 0) {
           totalSame++;
           alphabetList[getIndex(char)]--;
       }
   }
 
   return s1.length - totalSame + s2.length - totalSame;
}

function solve(): void {
    const testCases:string[] =  ["ayam bayam", "kursi bui", "ikan sapi", "hitam putih"];
   
    for(const testCase of testCases) {
        console.log(countTotalSteps(testCase));
    }
}

solve();
